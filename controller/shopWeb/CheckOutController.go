package shopWeb

import (
	"ShopWebGo/model"
	"ShopWebGo/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CheckOutController struct {
	BaseController
}

func (con CheckOutController) Checkout(c *gin.Context) {
	//1、获取购物车中选择的商品

	cartList := []model.Cart{}
	util.Cookie.Get(c, "cartList", &cartList)

	orderList := []model.Cart{}
	var allPrice float64
	var allNum int

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
			allNum += cartList[i].Num
		}
	}

	//2、获取当前用户的收货地址

	userId, _, ok := util.GetUserFromJWT(c)
	if !ok {
		c.Redirect(302, "/pass/login")
		return
	}
	addressList := []model.Address{}
	util.DB.Where("uid = ?", userId).Order("id desc").Find(&addressList)

	//3、生成签名
	orderSign := util.Md5(util.GetRandomNum())
	session := sessions.Default(c)
	session.Set("orderSign", orderSign)
	session.Save()

	//4、判断orderList数据是否存在
	if len(orderList) == 0 {
		c.Redirect(302, "/")
		return
	}

	con.Render(c, "shopWeb/buy/checkout.html", gin.H{
		"orderList":   orderList,
		"allPrice":    allPrice,
		"allNum":      allNum,
		"addressList": addressList,
		"orderSign":   orderSign,
	})

}

/*
提交订单执行结算

	1、获取用户信息 获取用户的收货地址信息
	2、获取购买商品的信息
	3、把订单信息放在订单表，把商品信息放在商品表
	4、删除购物车里面的选中数据
	5、跳转到支付页面
*/
func (con CheckOutController) DoCheckout(c *gin.Context) {
	//0、防止重复提交订单
	orderSignClient := c.PostForm("orderSign")
	session := sessions.Default(c)
	orderSignSession := session.Get("orderSign")
	orderSignServer, ok := orderSignSession.(string)
	if !ok {
		c.Redirect(302, "/")
		return
	}

	if orderSignClient != orderSignServer {
		c.Redirect(302, "/")
		return
	}
	session.Delete("orderSign")
	session.Save()

	// 1、获取用户信息 获取用户的收货地址信息
	userId, _, ok := util.GetUserFromJWT(c)
	if !ok {
		c.Redirect(302, "/pass/login")
		return
	}

	addressResult := []model.Address{}
	util.DB.Where("uid = ? AND default_address=1", userId).Find(&addressResult)
	if len(addressResult) == 0 {
		c.Redirect(302, "/buy/checkout")
		return
	}

	// 2、获取购买商品的信息
	cartList := []model.Cart{}
	util.Cookie.Get(c, "cartList", &cartList)
	orderList := []model.Cart{}
	var allPrice float64
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}
	// 3、事务包裹：把订单信息和商品信息原子写入
	order := model.Order{
		OrderId:     util.GetOrderId(),
		Uid:         userId,
		AllPrice:    allPrice,
		Phone:       addressResult[0].Phone,
		Name:        addressResult[0].Name,
		Address:     addressResult[0].Address,
		PayStatus:   0,
		PayType:     0,
		OrderStatus: 0,
		AddTime:     int(util.GetUnix()),
	}

	err := util.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		for i := 0; i < len(orderList); i++ {
			orderItem := model.OrderItem{
				OrderId:      order.Id,
				Uid:          userId,
				ProductTitle: orderList[i].Title,
				ProductId:    orderList[i].Id,
				ProductImg:   orderList[i].GoodsImg,
				ProductPrice: orderList[i].Price,
				ProductNum:   orderList[i].Num,
				GoodsVersion: orderList[i].GoodsVersion,
				GoodsColor:   orderList[i].GoodsColor,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		c.Redirect(302, "/buy/checkout")
		return
	}

	for i := 0; i < len(orderList); i++ {
		util.IndexOrderItem(util.OrderItemDoc{
			OrderId:      order.Id,
			Uid:          userId,
			ProductTitle: orderList[i].Title,
			ProductPrice: orderList[i].Price,
		})
	}

	// 4、删除购物车里面的选中数据
	noSelectCartList := []model.Cart{}
	for i := 0; i < len(cartList); i++ {
		if !cartList[i].Checked {
			noSelectCartList = append(noSelectCartList, cartList[i])
		}
	}
	util.Cookie.Set(c, "cartList", noSelectCartList)

	c.Redirect(302, "/buy/pay?orderId="+util.String(order.Id))
}

// 支付
func (con CheckOutController) Pay(c *gin.Context) {

	orderId, err := util.Int(c.Query("orderId"))
	if err != nil {
		c.Redirect(302, "/")
	}
	//获取用户信息
	userId, _, ok := util.GetUserFromJWT(c)
	if !ok {
		c.Redirect(302, "/pass/login")
		return
	}
	//获取订单信息
	order := model.Order{}
	util.DB.Where("id = ?", orderId).Find(&order)
	if order.Uid != userId {
		c.Redirect(302, "/")
		return
	}
	//获取订单对应的商品

	orderItems := []model.OrderItem{}
	util.DB.Where("order_id = ?", orderId).Find(&orderItems)

	con.Render(c, "shopWeb/buy/pay.html", gin.H{
		"order":      order,
		"orderItems": orderItems,
	})
}

func (con CheckOutController) DoPay(c *gin.Context) {
	orderId, err := util.Int(c.Query("orderId"))
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	userId, _, ok := util.GetUserFromJWT(c)
	if !ok {
		c.Redirect(302, "/pass/login")
		return
	}

	order := model.Order{}
	util.DB.Where("id = ? AND uid = ?", orderId, userId).Find(&order)
	if order.Id == 0 {
		c.Redirect(302, "/")
		return
	}

	if order.PayStatus == 1 {
		c.Redirect(302, "/user/order")
		return
	}

	util.DB.Model(&order).Updates(map[string]interface{}{
		"pay_status":   1,
		"order_status": 1,
	})

	c.Redirect(302, "/user/order")
}
