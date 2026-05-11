package shopWeb

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
	BaseController
}

func (con DefaultController) Index(c *gin.Context) {

	timeStart := time.Now().UnixNano()
	//1、获取顶部导航

	//2、获取轮播图数据
	focusList := []model.Focus{}
	if hasFocusList := util.CacheDb.Get("focusList", &focusList); !hasFocusList {
		util.DB.Where("status=1 AND focus_type=1").Find(&focusList)
		util.CacheDb.Set("focusList", focusList, 60*60)
	}

	//3、获取分类的数据

	//4、获取中间导航

	//手机
	phoneList := []model.Goods{}
	if hasPhoneList := util.CacheDb.Get("phoneList", &phoneList); !hasPhoneList {
		phoneList = util.GetGoodsByCategory(23, "best", 8)
		util.CacheDb.Set("phoneList", phoneList, 60*60)
	}

	//配件

	otherList := []model.Goods{}
	if hasOtherList := util.CacheDb.Get("otherList", &otherList); !hasOtherList {
		otherList = util.GetGoodsByCategory(9, "all", 1)
		util.CacheDb.Set("otherList", otherList, 60*60)
	}

	timeEnd := time.Now().UnixNano()

	fmt.Printf("执行时间：%v 毫秒", (timeEnd-timeStart)/1000000)

	con.Render(c, "shopWeb/index/index.html", gin.H{
		"focusList": focusList,
		"phoneList": phoneList,
		"otherList": otherList,
	})

}

// 将执行时间优化为1毫秒左右
