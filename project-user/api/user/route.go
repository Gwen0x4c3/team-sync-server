package user

import (
	"log"

	"github.com/Gwen0x4c3/team-sync-server/project-user/router"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// Route 从这里定义user模块的路由
func (ur *UserRouter) Route(r *gin.Engine) {
	handler := &HandlerUser{}
	r.GET("/project/login/getCaptcha", handler.getCaptcha)
}

func init() {
	log.Println("Init user router")
	router.Register(&UserRouter{})
}
