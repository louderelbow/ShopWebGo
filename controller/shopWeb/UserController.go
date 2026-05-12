package shopWeb

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"fmt"
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

	userId, _, ok := util.GetUserFromJWT(c)
	if !ok {
		c.Redirect(302, "/pass/login")
		return
	}

	keywords := c.Query("keywords")

	var orderIds []int
	if keywords != "" {
		orderIds = util.SearchOrderItems(userId, keywords)
		if orderIds != nil {
			fmt.Printf("[订单搜索] 使用 Elasticsearch | userId=%d keywords=%s 结果数=%d\n", userId, keywords, len(orderIds))
		} else {
			orderItemList := []model.OrderItem{}
			util.DB.Where("product_title LIKE ?", "%"+keywords+"%").Find(&orderItemList)
			for _, v := range orderItemList {
				orderIds = append(orderIds, v.OrderId)
			}
			fmt.Printf("[订单搜索] 降级 MySQL LIKE | userId=%d keywords=%s 结果数=%d\n", userId, keywords, len(orderIds))
		}
	}

	orderList := []model.Order{}
	listDB := util.DB.Where("uid = ?", userId)
	if keywords != "" {
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
	countDB := util.DB.Table("order").Where("uid = ?", userId)
	if keywords != "" {
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
	userId, _, ok := util.GetUserFromJWT(c)
	if !ok {
		c.Redirect(302, "/pass/login")
		return
	}
	order := []model.Order{}
	util.DB.Where("id=? And uid=?", id, userId).Preload("OrderItem").Find(&order)

	if len(order) == 0 {
		c.Redirect(302, "/user/order")
		return
	}
	var tpl = "shopWeb/user/order_info.html"
	con.Render(c, tpl, gin.H{
		"order": order[0],
	})
}
