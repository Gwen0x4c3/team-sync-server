package user

import (
	"context"
	"net/http"
	"time"

	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	"github.com/Gwen0x4c3/team-sync-server/project-common/errs"
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	loginServiceV1 "github.com/Gwen0x4c3/team-sync-server/project-user/pkg/service/login.service.v1"
	"github.com/gin-gonic/gin"
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
		code, msg := errs.ParseGrpcError(err)
		logs.LG.Error("GetCaptcha error: %v\n", err)
		c.JSON(http.StatusOK, result.Error(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success(response.Code))
}
