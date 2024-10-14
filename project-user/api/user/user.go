package user

import (
	"context"
	"fmt"
	"github.com/Gwen0x4c3/team-sync-server/project-user/pkg/constant"
	"github.com/Gwen0x4c3/team-sync-server/project-user/pkg/dao"
	"github.com/Gwen0x4c3/team-sync-server/project-user/pkg/model"
	"github.com/Gwen0x4c3/team-sync-server/project-user/pkg/repo"
	"log"
	"math/rand"
	"net/http"
	"time"

	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cache repo.Cache
}

func NewHandler() *Handler {
	user := Handler{
		cache: dao.Rc,
	}
	return &user
}

// getCaptcha 获取验证码
func (handler *Handler) getCaptcha(c *gin.Context) {
	log.Println("正在生成验证码")
	resp := &common.Result{}
	// 1. 获取参数
	mobile := c.PostForm("mobile")
	// 2. 校验参数
	if !common.VerifyMobile(mobile) {
		c.JSON(http.StatusOK, resp.Error(model.IllegalMobile, "手机号格式错误"))
		return
	}
	// 3. 生成验证码（1000-9999）
	// 随机生成一个4位数的验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	// 4. 调用短信平台（三方，可以放入协程，快速返回响应）
	go func() {
		// 模拟发送短信服务
		time.Sleep(2 * time.Second)
		log.Printf("已向手机号【%s】发送验证码：%s\n", mobile, code)

		// 5. 存储验证码
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		redisKey := constant.MakeRedisKey(constant.UserCaptchaKey, mobile)
		err := handler.cache.Put(ctx, constant.MakeRedisKey(redisKey, mobile), code, 2*time.Minute)
		if err != nil {
			c.JSON(http.StatusOK, resp.Error(http.StatusInternalServerError, "存储验证码失败"))
			log.Fatalf("存储验证码失败：%v\n", err)
		}
		log.Printf("将手机号【%s】的验证码%s存入缓存\n", mobile, code)
	}()
	c.JSON(http.StatusOK, resp.Success(code))
}
