package user

import (
	"context"
	"github.com/Gwen0x4c3/team-sync-server/project-api/pkg/model"
	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	loginServiceV1 "github.com/Gwen0x4c3/team-sync-server/project-user/pkg/service/login.service.v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (h *HandlerUser) GetCaptcha(c *gin.Context) {
	result := &common.Result{}
	mobile := c.PostForm("mobile")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := loginServiceClient.GetCaptcha(ctx, &loginServiceV1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		logs.LG.Error("GetCaptcha error: %v\n", err)
		c.JSON(http.StatusOK, result.Error(model.IllegalMobile, "获取验证码失败"))
		return
	}
	c.JSON(http.StatusOK, result.Success(response.Code))
}
