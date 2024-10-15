package user

import (
	"github.com/Gwen0x4c3/team-sync-server/project-api/router"
	"github.com/gin-gonic/gin"
	"log"
)

type Router struct {
}

var handler = New()

func (*Router) Route(r *gin.Engine) {
	InitRpcClient()
	r.GET("/project/user/getCaptcha", handler.GetCaptcha)
}

func init() {
	log.Println("Init user router")
	ro := &Router{}
	router.Register(ro)
}
