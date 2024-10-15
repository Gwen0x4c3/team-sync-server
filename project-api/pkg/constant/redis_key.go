package constant

import "fmt"

const (
	UserCaptchaKey = "user:captcha:%s"
)

func MakeRedisKey(key string, args ...interface{}) string {
	return fmt.Sprintf(key, args...)
}
