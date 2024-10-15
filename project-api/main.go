package main

import (
	_ "github.com/Gwen0x4c3/team-sync-server/project-api/api"
	"github.com/Gwen0x4c3/team-sync-server/project-api/config"
	"github.com/Gwen0x4c3/team-sync-server/project-api/router"
	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	"github.com/gin-gonic/gin"
)

func main() {
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
	logs.LG.Info("Init %s logger", config.Cfg.Server.Name)

	// 初始化gin
	r := gin.Default()
	r.Use(logs.GinLogger(), logs.GinRecovery(true))

	// 初始化路由
	router.InitRouter(r)
	common.Run(r, config.Cfg.Server.Name, config.Cfg.Server.Addr, nil)
}
