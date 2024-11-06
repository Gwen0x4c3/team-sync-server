package user

import (
	"log"

	"github.com/Gwen0x4c3/team-sync-server/project-api/router"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

var handler = New()

func (*Router) Route(r *gin.Engine) {
	InitRpcClient()
	r.POST("/project/user/getCaptcha", handler.GetCaptcha)
}

func init() {
	log.Println("Init user router")
	ro := &Router{}
	router.Register(ro)
}
