package middlewares

import (
	"ShopWebGo/util"

	"github.com/gin-gonic/gin"
)

func InitUserAuthMiddleware(c *gin.Context) {
	userId, _, ok := util.GetUserFromJWT(c)
	if !ok || userId == 0 {
		c.Redirect(302, "/pass/login")
		return
	}
}
