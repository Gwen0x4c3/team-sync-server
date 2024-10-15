package main

import (
	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	_ "github.com/Gwen0x4c3/team-sync-server/project-user/api"
	"github.com/Gwen0x4c3/team-sync-server/project-user/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	// 初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "D:\\Learning\\Projects\\GO\\2.go-team-sync\\team-sync-server\\logs\\debug\\project-debug.log",
		InfoFileName:  "D:\\Learning\\Projects\\GO\\2.go-team-sync\\team-sync-server\\logs\\info\\project-info.log",
		WarnFileName:  "D:\\Learning\\Projects\\GO\\2.go-team-sync\\team-sync-server\\logs\\error\\project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}

	// 初始化路由
	router.InitRouter(r)
	common.Run(r, "project-user", ":80")
}
