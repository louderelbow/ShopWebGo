package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (con BaseController) success(c *gin.Context, message string, redirectUrl string) {

	//传入gin.H为了让前端渲染 前后端分离项目需要放到返回值结构体里
	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (con BaseController) error(c *gin.Context, message string, redirectUrl string) {

	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
