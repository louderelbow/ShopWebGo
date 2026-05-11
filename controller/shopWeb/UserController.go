package shopWeb

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"math"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	// c.String(http.StatusOK, "首页")
	var tpl = "shopWeb/user/welcome.html"
	con.Render(c, tpl, gin.H{})
}
func (con UserController) OrderList(c *gin.Context) {
	page, _ := util.Int(c.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize := 2

	user := model.User{}
	util.Cookie.Get(c, "userinfo", &user)

	keywords := c.Query("keywords")

	orderList := []model.Order{}
	listDB := util.DB.Where("uid = ?", user.Id)

	if keywords != "" {
		orderItemList := []model.OrderItem{}
		util.DB.Where("product_title LIKE ?", "%"+keywords+"%").Find(&orderItemList)
		var orderIds []int
		for _, v := range orderItemList {
			orderIds = append(orderIds, v.OrderId)
		}
		if len(orderIds) > 0 {
			listDB = listDB.Where("id IN ?", orderIds)
		} else {
			listDB = listDB.Where("id IN ?", []int{0})
		}
	}

	orderStatus, statusErr := util.Int(c.Query("orderStatus"))
	if statusErr == nil && orderStatus >= 0 {
		listDB = listDB.Where("order_status = ?", orderStatus)
	} else {
		orderStatus = -1
	}

	listDB.Preload("OrderItem").Order("add_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&orderList)

	var count int64
	countDB := util.DB.Table("order").Where("uid = ?", user.Id)
	if keywords != "" {
		orderItemList := []model.OrderItem{}
		util.DB.Where("product_title LIKE ?", "%"+keywords+"%").Find(&orderItemList)
		var orderIds []int
		for _, v := range orderItemList {
			orderIds = append(orderIds, v.OrderId)
		}
		if len(orderIds) > 0 {
			countDB = countDB.Where("id IN ?", orderIds)
		} else {
			countDB = countDB.Where("id IN ?", []int{0})
		}
	}
	if statusErr == nil && orderStatus >= 0 {
		countDB = countDB.Where("order_status = ?", orderStatus)
	}
	countDB.Count(&count)

	var tpl = "shopWeb/user/order.html"
	con.Render(c, tpl, gin.H{
		"order":       orderList,
		"page":        page,
		"keywords":    keywords,
		"orderStatus": orderStatus,
		"totalPages":  math.Ceil(float64(count) / float64(pageSize)),
	})
}
func (con UserController) OrderInfo(c *gin.Context) {

	id, err := util.Int(c.Query("id"))
	if err != nil {
		c.Redirect(302, "/user/order")
	}
	user := model.User{}
	util.Cookie.Get(c, "userinfo", &user)
	order := []model.Order{}
	util.DB.Where("id=? And uid=?", id, user.Id).Preload("OrderItem").Find(&order)

	if len(order) == 0 {
		c.Redirect(302, "/user/order")
		return
	}
	var tpl = "shopWeb/user/order_info.html"
	con.Render(c, tpl, gin.H{
		"order": order[0],
	})
}
