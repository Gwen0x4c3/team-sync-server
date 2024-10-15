package user

import (
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	"github.com/Gwen0x4c3/team-sync-server/project-user/router"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// Route 从这里定义user模块的路由
func (ur *UserRouter) Route(r *gin.Engine) {
	handler := NewHandler()
	// 验证码
	r.POST("/project/login/getCaptcha", handler.getCaptcha)
}

func init() {
	logs.Log.Info("Init user router")
	router.Register(&UserRouter{})
}
