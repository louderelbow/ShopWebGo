package shopWeb

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"fmt"

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

	//5、读取上次提交的错误消息（如果有）
	checkoutError := ""
	if errMsg := session.Get("checkoutError"); errMsg != nil {
		if errStr, ok := errMsg.(string); ok {
			checkoutError = errStr
		}
		session.Delete("checkoutError")
		session.Save()
	}

	con.Render(c, "shopWeb/buy/checkout.html", gin.H{
		"orderList":     orderList,
		"allPrice":      allPrice,
		"allNum":        allNum,
		"addressList":   addressList,
		"orderSign":     orderSign,
		"checkoutError": checkoutError,
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
	fmt.Println("[DoCheckout] ========== 开始提交订单 ==========")

	//0、防止重复提交订单
	orderSignClient := c.PostForm("orderSign")
	session := sessions.Default(c)
	orderSignSession := session.Get("orderSign")
	orderSignServer, ok := orderSignSession.(string)
	fmt.Printf("[DoCheckout] orderSign client=%s server=%s ok=%v\n", orderSignClient, orderSignServer, ok)
	if !ok {
		fmt.Println("[DoCheckout] ❌ session中无orderSign")
		c.Redirect(302, "/")
		return
	}

	if orderSignClient != orderSignServer {
		fmt.Println("[DoCheckout] ❌ orderSign不匹配（重复提交或非法请求）")
		c.Redirect(302, "/")
		return
	}
	session.Delete("orderSign")
	session.Save()
	fmt.Println("[DoCheckout] ✅ orderSign校验通过")

	// 1、获取用户信息 获取用户的收货地址信息
	userId, _, ok := util.GetUserFromJWT(c)
	if !ok {
		fmt.Println("[DoCheckout] ❌ 未登录，跳转登录页")
		c.Redirect(302, "/pass/login")
		return
	}
	fmt.Printf("[DoCheckout] userId=%d\n", userId)

	addressResult := []model.Address{}
	util.DB.Where("uid = ? AND default_address=1", userId).Find(&addressResult)
	if len(addressResult) == 0 {
		fmt.Printf("[DoCheckout] ❌ 用户%d 没有默认收货地址\n", userId)
		session.Set("checkoutError", "请先设置默认收货地址")
		session.Save()
		c.Redirect(302, "/buy/checkout")
		return
	}
	fmt.Printf("[DoCheckout] ✅ 默认地址: name=%s phone=%s address=%s\n", addressResult[0].Name, addressResult[0].Phone, addressResult[0].Address)

	// 2、获取购买商品的信息
	cartList := []model.Cart{}
	util.Cookie.Get(c, "cartList", &cartList)
	fmt.Printf("[DoCheckout] 购物车商品总数=%d\n", len(cartList))

	orderList := []model.Cart{}
	var allPrice float64
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
			fmt.Printf("[DoCheckout] 📦 选中商品: id=%d title=%s price=%.2f num=%d\n",
				cartList[i].Id, cartList[i].Title, cartList[i].Price, cartList[i].Num)
		}
	}
	fmt.Printf("[DoCheckout] 订单商品数=%d 总价=%.2f\n", len(orderList), allPrice)

	if len(orderList) == 0 {
		fmt.Println("[DoCheckout] ❌ 没有选中商品")
		c.Redirect(302, "/")
		return
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
	fmt.Printf("[DoCheckout] 📤 准备创建订单: OrderId=%s Uid=%d AllPrice=%.2f\n", order.OrderId, order.Uid, order.AllPrice)

	err := util.DB.Transaction(func(tx *gorm.DB) error {
		// 扣库存：逐商品检查库存并扣减
		for i := 0; i < len(orderList); i++ {
			fmt.Printf("[DoCheckout] 🔒 扣库存: goodsId=%d title=%s need=%d\n",
				orderList[i].Id, orderList[i].Title, orderList[i].Num)

			result := tx.Model(&model.Goods{}).
				Where("id = ? AND goods_number >= ? AND status = 1", orderList[i].Id, orderList[i].Num).
				UpdateColumn("goods_number", gorm.Expr("goods_number - ?", orderList[i].Num))
			if result.RowsAffected == 0 {
				fmt.Printf("[DoCheckout] ❌ 扣库存失败: goodsId=%d title=%s 库存不足或已下架\n",
					orderList[i].Id, orderList[i].Title)
				return fmt.Errorf("商品 %s 库存不足或已下架", orderList[i].Title)
			}
			fmt.Printf("[DoCheckout] ✅ 扣库存成功: goodsId=%d\n", orderList[i].Id)
		}

		fmt.Println("[DoCheckout] 📤 写入order表...")
		if err := tx.Create(&order).Error; err != nil {
			fmt.Printf("[DoCheckout] ❌ 创建订单失败: %v\n", err)
			return err
		}
		fmt.Printf("[DoCheckout] ✅ order写入成功 OrderDBId=%d\n", order.Id)

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
			fmt.Printf("[DoCheckout] 📤 写入orderItem: product=%s price=%.2f num=%d\n",
				orderItem.ProductTitle, orderItem.ProductPrice, orderItem.ProductNum)
			if err := tx.Create(&orderItem).Error; err != nil {
				fmt.Printf("[DoCheckout] ❌ 创建订单商品失败: %v\n", err)
				return err
			}
			fmt.Printf("[DoCheckout] ✅ orderItem写入成功 id=%d\n", orderItem.Id)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("[DoCheckout] ❌ 事务失败回滚: %v\n", err)
		session.Set("checkoutError", err.Error())
		session.Save()
		c.Redirect(302, "/buy/checkout")
		return
	}
	fmt.Println("[DoCheckout] ✅ 事务提交成功")

	// ES 索引（非关键路径）
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
	fmt.Printf("[DoCheckout] 购物车清理: 保留%d个未选中商品\n", len(noSelectCartList))

	fmt.Printf("[DoCheckout] ✅ 订单提交完成，跳转支付页 orderId=%d\n", order.Id)
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
