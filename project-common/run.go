package common

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/graceful"
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine, srvName, srvAddr string) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	router, err := graceful.New(r, graceful.WithAddr(srvAddr))

	if err != nil {
		log.Fatalf("Failed to create server: %s\n", err)
		panic(err)
	}
	defer router.Close()

	go func() {
		if err := router.RunWithContext(ctx); err != nil && err != context.Canceled {
			log.Fatalf("Failed to start server: %s\n", err)
		}
	}()

	fmt.Printf("%s is running on %s\n", srvName, srvAddr)
	<-ctx.Done()

	fmt.Println("Server has been stopped")
}
