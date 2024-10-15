package main

import (
	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	common.Run(r, "project-api", ":9999", nil)
}
