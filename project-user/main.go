package main

import (
	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	_ "github.com/Gwen0x4c3/team-sync-server/project-user/api"
	"github.com/Gwen0x4c3/team-sync-server/project-user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	common.Run(r, "project-user", ":10001")
}
