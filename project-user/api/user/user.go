package user

import (
	"log"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct{}

// getCaptcha 获取验证码
func (handler *HandlerUser) getCaptcha(ctx *gin.Context) {
	log.Println("Handle the getCaptcha request")
	ctx.JSON(200, gin.H{
		"message": "getCaptcha",
	})
}
