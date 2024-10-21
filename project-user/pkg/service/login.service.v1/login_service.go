package login_service_v1

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Gwen0x4c3/team-sync-server/project-api/pkg/model"
	common "github.com/Gwen0x4c3/team-sync-server/project-common"
	"github.com/Gwen0x4c3/team-sync-server/project-common/errs"
	"github.com/Gwen0x4c3/team-sync-server/project-common/logs"
	"github.com/Gwen0x4c3/team-sync-server/project-user/pkg/constant"
	"github.com/Gwen0x4c3/team-sync-server/project-user/pkg/dao"
	"github.com/Gwen0x4c3/team-sync-server/project-user/pkg/repo"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func NewLoginService() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (service *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	// 1. 获取参数
	mobile := msg.Mobile
	// 2. 校验参数
	if !common.VerifyMobile(mobile) {
		// TODO 以后改成带错误码的返回 model.IllegalMobile
		return nil, errs.GrpcError(model.IllegalMobile)
	}
	// 3. 生成验证码（1000-9999）
	// 随机生成一个4位数的验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	// 4. 调用短信平台（三方，可以放入协程，快速返回响应）
	go func() {
		// 模拟发送短信服务
		time.Sleep(1 * time.Second)
		logs.LG.Info("已向手机号【%s】发送验证码：%s\n", mobile, code)

		// 5. 存储验证码
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		redisKey := constant.MakeRedisKey(constant.UserCaptchaKey, mobile)
		err := service.cache.Put(ctx, constant.MakeRedisKey(redisKey, mobile), code, 2*time.Minute)
		if err != nil {
			logs.LG.Error("存储验证码失败：%v\n", err)
		}
		logs.LG.Info("将手机号【%s】的验证码%s存入缓存\n", mobile, code)
	}()
	return &CaptchaResponse{Code: code}, nil
}
