package main

import (
	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	_ "github.com/Gwen0x4c3/team-sync-server/project-user/api"
	"github.com/Gwen0x4c3/team-sync-server/project-user/config"
	"github.com/Gwen0x4c3/team-sync-server/project-user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化日志
	lc := &logs.LogConfig{
		DebugFileName: config.Cfg.Zap.DebugFileName,
		InfoFileName:  config.Cfg.Zap.InfoFileName,
		WarnFileName:  config.Cfg.Zap.WarnFileName,
		MaxSize:       config.Cfg.Zap.MaxSize,
		MaxAge:        config.Cfg.Zap.MaxAge,
		MaxBackups:    config.Cfg.Zap.MaxBackups,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		panic(err)
	}
	logs.LG.Info("Init project-user logger")
	r.Use(logs.GinLogger(), logs.GinRecovery(true))

	// 初始化路由
	router.InitRouter(r)
	common.Run(r, config.Cfg.Server.Name, config.Cfg.Server.Addr)
}
