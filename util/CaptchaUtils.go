package util

import (
	"fmt"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 创建store
// var store = base64Captcha.DefaultMemStore
// 配置RedisStore  RedisStore实现base64Captcha.Store接口
// 由于要将验证码存入Redis 所以这里实现的是我们自定义的Store
var store = base64Captcha.DefaultMemStore

// 获取验证码
func MakeCaptcha(height int, width int, length int) (string, string, error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          height,
		Width:           width,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          length,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()
	return id, b64s, err

}

// 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	fmt.Println(id, VerifyValue)
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
