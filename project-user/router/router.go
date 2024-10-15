package router

import (
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	"github.com/Gwen0x4c3/team-sync-server/project-user/config"
	loginServiceV1 "github.com/Gwen0x4c3/team-sync-server/project-user/pkg/service/login.service.v1"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Router interface {
	Route(r *gin.Engine)
}

// 注册路由第一种方法使用的结构体与方法
// type RegisterRouter struct {
// }

// func New() *RegisterRouter {
// 	return &RegisterRouter{}
// }

// func (rg *RegisterRouter) Route(rt Router, r *gin.Engine) {
// 	rt.Route(r)
// }

var routers []Router

func InitRouter(r *gin.Engine) {
	// 注册路由的第一种方法
	// rg := New()
	// rg.Route(&user.UserRouter{}, r)

	// 第二种方法，需要import _ "github.com/Gwen0x4c3/team-sync-server/project-user/api"来让各模块的route中的init函数执行
	for _, rt := range routers {
		rt.Route(r)
	}
}

func Register(rts ...Router) {
	routers = append(routers, rts...)
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.Cfg.Grpc.Addr,
		RegisterFunc: func(server *grpc.Server) {
			loginServiceV1.RegisterLoginServiceServer(server, loginServiceV1.NewLoginService())
		},
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)
	listen, err := net.Listen("tcp", c.Addr)
	if err != nil {
		logs.LG.Error("failed to listen: %v", zap.Error(err))
	}
	go func() {
		if err := s.Serve(listen); err != nil {
			logs.LG.Error("failed to serve: %v", zap.Error(err))
		}
	}()
	return s
}
