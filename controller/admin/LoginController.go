package admin

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(c *gin.Context) {
	//获取验证码ID
	captchaId := c.PostForm("captchaId")
	//获取用户名密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	//获取验证码输入值
	verifyValue := c.PostForm("verifyValue")
	//将输入内容转化为小写
	verifyValue = strings.ToLower(verifyValue)
	//验证是否正确
	if flag := util.VerifyCaptcha(captchaId, verifyValue); flag {
		//2、查询数据库 判断用户以及密码是否存在
		userinfoList := []model.Manager{}
		password = util.Md5(password)

		util.DB.Where("username=? AND password=?", username, password).Find(&userinfoList)

		if len(userinfoList) > 0 {
			//3、执行登录 保存用户信息 执行跳转
			session := sessions.Default(c)
			//注意：session.Set没法直接保存结构体对应的切片 把结构体转换成json字符串
			userinfoSlice, _ := json.Marshal(userinfoList)
			fmt.Println(string(userinfoSlice))
			session.Set("userinfo", string(userinfoSlice))
			session.Save()
			con.success(c, "登录成功", "/admin")

		} else {
			con.error(c, "用户名或者密码错误", "/admin/login")
		}

	} else {
		con.error(c, "验证码验证失败", "/admin/login")
	}
}

func (LoginController) Captcha(c *gin.Context) {
	id, b64s, err := util.MakeCaptcha(35, 100, 2)

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (con LoginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.success(c, "退出登录成功", "/admin/login")
}
