package admin

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	focusList := []model.Focus{}
	util.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})

}
func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
func (con FocusController) DoAdd(c *gin.Context) {

	title := c.PostForm("title")
	focusType, err1 := util.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err2 := util.Int(c.PostForm("sort"))
	status, err3 := util.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.error(c, "非法请求", "/admin/focus/add")
	}
	if err2 != nil {
		con.error(c, "请输入正确的排序值", "/admin/focus/add")
	}
	//上传文件
	focusImgSrc, err4 := util.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}

	focus := model.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(util.GetUnix()),
	}
	err5 := util.DB.Create(&focus).Error
	if err5 != nil {
		con.error(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		con.success(c, "增加轮播图成功", "/admin/focus")
	}

}

func (con FocusController) Edit(c *gin.Context) {
	id, err1 := util.Int(c.Query("id"))
	if err1 != nil {
		con.error(c, "传入参数错误", "/admin/focus")
		return
	}
	focus := model.Focus{Id: id}
	util.DB.Find(&focus)
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}
func (con FocusController) DoEdit(c *gin.Context) {
	id, err1 := util.Int(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := util.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := util.Int(c.PostForm("sort"))
	status, err4 := util.Int(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.error(c, "非法请求", "/admin/focus")
	}
	if err3 != nil {
		con.error(c, "请输入正确的排序值", "/admin/focus/edit?id="+util.String(id))
	}
	//上传文件
	focusImg, _ := util.UploadImg(c, "focus_img")

	focus := model.Focus{Id: id}
	util.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}
	err5 := util.DB.Save(&focus).Error
	if err5 != nil {
		con.error(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+util.String(id))
	} else {
		con.success(c, "增加轮播图成功", "/admin/focus")
	}
}

func (con FocusController) Delete(c *gin.Context) {
	id, err := util.Int(c.Query("id"))
	if err != nil {
		con.error(c, "传入数据错误", "/admin/focus")
	} else {
		focus := model.Focus{Id: id}
		util.DB.Delete(&focus)
		//根据自己的需要 要不要删除图片
		// os.Remove("static/upload/20260428/1631694117.jpg")
		con.success(c, "删除数据成功", "/admin/focus")
	}
}
