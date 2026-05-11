package admin

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NavController struct {
	BaseController
}

func (con NavController) Index(c *gin.Context) {
	//当前页
	page, _ := util.Int(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//每页显示的数量
	pageSize := 8
	//获取数据
	navList := []model.Nav{}
	util.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&navList)

	//获取总数量
	var count int64
	util.DB.Table("nav").Count(&count)
	c.HTML(http.StatusOK, "admin/nav/index.html", gin.H{
		"navList": navList,
		//注意float64类型
		"totalPages": math.Ceil(float64(count) / float64(pageSize)),
		"page":       page,
	})
}
func (con NavController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/nav/add.html", gin.H{})
}
func (con NavController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := util.Int(c.PostForm("position"))
	isOpennew, _ := util.Int(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := util.Int(c.PostForm("sort"))
	status, _ := util.Int(c.PostForm("status"))
	if title == "" {
		con.error(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := model.Nav{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(util.GetUnix()),
	}
	err := util.DB.Create(&nav).Error
	if err != nil {
		con.error(c, "增加导航失败 请重试", "/admin/nav/add")
	} else {
		con.success(c, "增加导航成功", "/admin/nav")
	}

}
func (con NavController) Edit(c *gin.Context) {
	id, err := util.Int(c.Query("id"))
	if err != nil {
		con.error(c, "传入数据错误", "/admin/nav")
	} else {
		nav := model.Nav{Id: id}
		util.DB.Find(&nav)
		fmt.Println(nav)
		c.HTML(http.StatusOK, "admin/nav/edit.html", gin.H{
			"nav": nav,
		})
	}

}
func (con NavController) DoEdit(c *gin.Context) {

	id, err1 := util.Int(c.PostForm("id"))
	if err1 != nil {
		con.error(c, "传入数据错误", "/admin/nav")
		return
	}

	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := util.Int(c.PostForm("position"))
	isOpennew, _ := util.Int(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := util.Int(c.PostForm("sort"))
	status, _ := util.Int(c.PostForm("status"))
	if title == "" {
		con.error(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := model.Nav{Id: id}
	util.DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status
	err2 := util.DB.Save(&nav).Error
	if err2 != nil {
		con.error(c, "修改数据失败", "/admin/nav/edit?id="+util.String(id))
	} else {
		con.success(c, "修改数据成功", "/admin/nav")
	}

}
func (con NavController) Delete(c *gin.Context) {
	id, err := util.Int(c.Query("id"))
	if err != nil {
		con.error(c, "传入数据错误", "/admin/nav")
	} else {
		nav := model.Nav{Id: id}
		util.DB.Delete(&nav)
		con.success(c, "删除数据成功", "/admin/nav")
	}
}
