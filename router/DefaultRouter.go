package router

import (
	"ShopWebGo/controller/shopWeb"
	"ShopWebGo/util/middlewares"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", shopWeb.DefaultController{}.Index)

		defaultRouters.GET("/category:id", shopWeb.ProductController{}.Category)
		defaultRouters.GET("/detail", shopWeb.ProductController{}.Detail)
		defaultRouters.GET("/product/getImgList", shopWeb.ProductController{}.GetImgList)

		defaultRouters.GET("/cart", shopWeb.CartController{}.Get)
		defaultRouters.GET("/cart/addCart", shopWeb.CartController{}.AddCart)

		defaultRouters.GET("/cart/successTip", shopWeb.CartController{}.AddCartSuccess)

		defaultRouters.GET("/cart/decCart", shopWeb.CartController{}.DecCart)
		defaultRouters.GET("/cart/incCart", shopWeb.CartController{}.IncCart)

		defaultRouters.GET("/cart/changeOneCart", shopWeb.CartController{}.ChangeOneCart)
		defaultRouters.GET("/cart/changeAllCart", shopWeb.CartController{}.ChangeAllCart)
		defaultRouters.GET("/cart/delCart", shopWeb.CartController{}.DelCart)

		defaultRouters.GET("/pass/login", shopWeb.PassController{}.Login)
		defaultRouters.GET("/pass/captcha", shopWeb.PassController{}.Captcha)

		defaultRouters.GET("/pass/registerStep1", shopWeb.PassController{}.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", shopWeb.PassController{}.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", shopWeb.PassController{}.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", shopWeb.PassController{}.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", shopWeb.PassController{}.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", shopWeb.PassController{}.DoRegister)
		defaultRouters.POST("/pass/doLogin", shopWeb.PassController{}.DoLogin)
		defaultRouters.GET("/pass/loginOut", shopWeb.PassController{}.LoginOut)

		defaultRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, shopWeb.CheckOutController{}.Checkout)
		defaultRouters.POST("/buy/doCheckout", middlewares.InitUserAuthMiddleware, shopWeb.CheckOutController{}.DoCheckout)
		defaultRouters.GET("/buy/pay", middlewares.InitUserAuthMiddleware, shopWeb.CheckOutController{}.Pay)
		defaultRouters.GET("/buy/doPay", middlewares.InitUserAuthMiddleware, shopWeb.CheckOutController{}.DoPay)

		defaultRouters.POST("/address/addAddress", middlewares.InitUserAuthMiddleware, shopWeb.AddressController{}.AddAddress)
		defaultRouters.POST("/address/editAddress", middlewares.InitUserAuthMiddleware, shopWeb.AddressController{}.EditAddress)
		defaultRouters.GET("/address/changeDefaultAddress", middlewares.InitUserAuthMiddleware, shopWeb.AddressController{}.ChangeDefaultAddress)
		defaultRouters.GET("/address/getOneAddressList", middlewares.InitUserAuthMiddleware, shopWeb.AddressController{}.GetOneAddressList)

		defaultRouters.GET("/user", middlewares.InitUserAuthMiddleware, shopWeb.UserController{}.Index)
		defaultRouters.GET("/user/order", middlewares.InitUserAuthMiddleware, shopWeb.UserController{}.OrderList)
		defaultRouters.GET("/user/orderinfo", middlewares.InitUserAuthMiddleware, shopWeb.UserController{}.OrderInfo)
	}
}
