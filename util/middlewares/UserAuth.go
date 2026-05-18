package middlewares

import (
	"ShopWebGo/util"

	"github.com/gin-gonic/gin"
)

// InitUserAuthMiddleware 用户认证中间件
func InitUserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _, ok := util.GetUserFromJWT(c)
		if !ok || userId == 0 {
			c.Redirect(302, "/pass/login")
			c.Abort()
			return
		}

		// 验证通过，放行
		c.Next()
	}
}