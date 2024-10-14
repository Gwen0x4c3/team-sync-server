package router

import (
	"github.com/gin-gonic/gin"
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
