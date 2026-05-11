package admin

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {

	managerList := []model.Manager{}
	util.DB.Preload("Role").Find(&managerList)

	fmt.Printf("%#v", managerList)

	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})

}
func (con ManagerController) Add(c *gin.Context) {
	//获取所有的角色
	roleList := []model.Role{}
	util.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}
func (con ManagerController) DoAdd(c *gin.Context) {

	roleId, err1 := util.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.error(c, "传入数据错误", "/admin/manager/add")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	//用户名和密码长度是否合法
	if len(username) < 2 || len(password) < 6 {
		con.error(c, "用户名或者密码的长度不合法", "/admin/manager/add")
		return
	}

	//判断管理是否存在
	managerList := []model.Manager{}
	util.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.error(c, "此管理员已存在", "/admin/manager/add")
		return
	}
	//执行增加管理员
	manager := model.Manager{
		Username: username,
		Password: util.Md5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   roleId,
		Status:   1,
		AddTime:  int(util.GetUnix()),
	}
	err2 := util.DB.Create(&manager).Error
	if err2 != nil {
		con.error(c, "增加管理员失败", "/admin/manager/add")
		return
	}

	con.success(c, "增加管理员成功", "/admin/manager")
}

func (con ManagerController) Edit(c *gin.Context) {

	//获取管理员
	id, err := util.Int(c.Query("id"))
	if err != nil {
		con.error(c, "传入数据错误", "/admin/manager")
		return
	}
	manager := model.Manager{Id: id}
	util.DB.Find(&manager)

	//获取所有的角色
	roleList := []model.Role{}
	util.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err1 := util.Int(c.PostForm("id"))
	if err1 != nil {
		con.error(c, "传入数据错误", "/admin/manager")
		return
	}
	roleId, err2 := util.Int(c.PostForm("role_id"))
	if err2 != nil {
		con.error(c, "传入数据错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")

	if len(mobile) > 11 {
		con.error(c, "mobile长度不合法", "/admin/manager/edit?id="+util.String(id))
		return
	}

	//执行修改
	manager := model.Manager{Id: id}
	util.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId
	manager.AddTime = int(util.GetUnix())

	//注意：判断密码是否为空 为空表示不修改密码 不为空表示修改密码

	if password != "" {
		//判断密码长度是否合法
		if len(password) < 6 {
			con.error(c, "密码的长度不合法 密码长度不能小于6位", "/admin/manager/edit?id="+util.String(id))
			return
		}
		manager.Password = util.Md5(password)
	}
	err3 := util.DB.Save(&manager).Error
	if err3 != nil {
		con.error(c, "修改数据失败", "/admin/manager/edit?id="+util.String(id))
		return
	}
	con.success(c, "修改数据成功", "/admin/manager")
}

func (con ManagerController) Delete(c *gin.Context) {
	id, err := util.Int(c.Query("id"))
	if err != nil {
		con.error(c, "传入数据错误", "/admin/manager")
	} else {
		manager := model.Manager{Id: id}
		util.DB.Delete(&manager)
		con.success(c, "删除数据成功", "/admin/manager")
	}
}
