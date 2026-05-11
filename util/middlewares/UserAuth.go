package middlewares

import (
	"ShopWebGo/model"
	"ShopWebGo/util"

	"github.com/gin-gonic/gin"
)

func InitUserAuthMiddleware(c *gin.Context) {
	//判断用户有没有登录
	user := model.User{}
	isLogin := util.Cookie.Get(c, "userinfo", &user)
	if !isLogin || len(user.Phone) != 11 {
		c.Redirect(302, "/pass/login")
		return
	}

}
