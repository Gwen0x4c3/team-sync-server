package common

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"

	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine, srvName, srvAddr string, stop func()) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	router, err := graceful.New(r, graceful.WithAddr(srvAddr))

	if err != nil {
		logs.LG.Error("Failed to create server: %s\n", err)
		panic(err)
	}
	defer router.Close()

	go func() {
		if err := router.RunWithContext(ctx); err != nil && err != context.Canceled {
			logs.LG.Error("Failed to run server: %s\n", err)
		}
	}()

	logs.LG.Info("%s server is running at %s\n", srvName, srvAddr)
	<-ctx.Done()

	if stop != nil {
		stop()
	}

	logs.LG.Info("Server has been stopped")
}
