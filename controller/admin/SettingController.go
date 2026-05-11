package admin

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type SettingController struct {
	BaseController
}

func (con SettingController) Index(c *gin.Context) {
	setting := model.Setting{}
	util.DB.First(&setting)
	c.HTML(http.StatusOK, "admin/setting/index.html", gin.H{
		"setting": setting,
	})
}

func (con SettingController) DoEdit(c *gin.Context) {
	setting := model.Setting{Id: 1}
	util.DB.Find(&setting)

	// 1. 绑定表单数据
	if err := c.ShouldBind(&setting); err != nil {
		con.error(c, "修改数据失败,请重试", "/admin/setting")
		return
	}

	// --------------- 协程并发上传图片（核心优化）---------------
	var wg sync.WaitGroup
	wg.Add(2) // 2个上传任务

	// 拷贝上下文（协程必用，防止请求结束上下文销毁）
	ctx := c.Copy()

	// 1. 协程1：上传 Logo
	go func() {
		defer wg.Done()
		siteLogo, err := util.UploadImg(ctx, "site_logo")
		if len(siteLogo) > 0 && err == nil {
			setting.SiteLogo = siteLogo
		}
	}()

	// 2. 协程2：上传默认商品图片
	go func() {
		defer wg.Done()
		noPicture, err := util.UploadImg(ctx, "no_picture")
		if len(noPicture) > 0 && err == nil {
			setting.NoPicture = noPicture
		}
	}()

	// 等待所有上传协程执行完毕
	wg.Wait()
	// ---------------------------------------------------------

	// 保存数据到数据库
	err3 := util.DB.Save(&setting).Error
	if err3 != nil {
		con.error(c, "修改数据失败", "/admin/setting")
		return
	}

	con.success(c, "修改数据成功", "/admin/setting")
}
