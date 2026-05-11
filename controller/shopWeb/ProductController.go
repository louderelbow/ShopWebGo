package shopWeb

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	BaseController
}

func (con ProductController) Category(c *gin.Context) {

	//分类id
	cateId, _ := util.Int(c.Param("id"))
	//当前页
	page, _ := util.Int(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 5
	//获取当前分类
	currentCate := model.GoodsCate{}
	util.DB.Where("id=?", cateId).Find(&currentCate)
	subCate := []model.GoodsCate{}
	var tempSlice []int
	if currentCate.Pid == 0 {
		//获取二级分类
		util.DB.Where("pid=?", currentCate.Id).Find(&subCate)
		for i := 0; i < len(subCate); i++ {
			tempSlice = append(tempSlice, subCate[i].Id)
		}
	} else {
		//兄弟分类
		util.DB.Where("pid=?", currentCate.Pid).Find(&subCate)
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in ?"
	goodsList := []model.Goods{}
	util.DB.Where(where, tempSlice).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsList)

	//获取总数量
	var count int64
	util.DB.Where(where, tempSlice).Table("goods").Count(&count)

	// 控制渲染的模板
	tpl := "shopWeb/product/list.html"
	if currentCate.Template != "" {
		tpl = currentCate.Template
	}

	con.Render(c, tpl, gin.H{
		"page":        page,
		"goodsList":   goodsList,
		"subCate":     subCate,
		"currentCate": currentCate,
		"totalPages":  math.Ceil(float64(count) / float64(pageSize)),
	})

}

func (con ProductController) Detail(c *gin.Context) {

	id, err := util.Int(c.Query("id"))

	if err != nil {
		c.Redirect(302, "/")
		c.Abort()
	}

	//1、获取商品信息（带缓存）
	goods := model.Goods{}
	goodsCacheKey := fmt.Sprintf("goods_detail_%d", id)
	if hasGoods := util.CacheDb.Get(goodsCacheKey, &goods); !hasGoods {
		util.DB.Find(&goods, id)
		util.CacheDb.Set(goodsCacheKey, goods, 3600)
	}

	//2、获取关联商品  RelationGoods（带缓存）
	relationGoods := []model.Goods{}
	goods.RelationGoods = strings.ReplaceAll(goods.RelationGoods, "，", ",")
	relationIds := strings.Split(goods.RelationGoods, ",")
	relationCacheKey := fmt.Sprintf("goods_relation_%d", id)
	if hasRelation := util.CacheDb.Get(relationCacheKey, &relationGoods); !hasRelation {
		if len(relationIds) > 0 && relationIds[0] != "" {
			util.DB.Where("id in ?", relationIds).Select("id,title,price,goods_version").Find(&relationGoods)
		}
		util.CacheDb.Set(relationCacheKey, relationGoods, 3600)
	}

	//3、获取关联赠品 GoodsGift（带缓存）
	goodsGift := []model.Goods{}
	goods.GoodsGift = strings.ReplaceAll(goods.GoodsGift, "，", ",")
	giftIds := strings.Split(goods.GoodsGift, ",")
	giftCacheKey := fmt.Sprintf("goods_gift_%d", id)
	if hasGift := util.CacheDb.Get(giftCacheKey, &goodsGift); !hasGift {
		if len(giftIds) > 0 && giftIds[0] != "" {
			util.DB.Where("id in ?", giftIds).Select("id,title,price,goods_version").Find(&goodsGift)
		}
		util.CacheDb.Set(giftCacheKey, goodsGift, 3600)
	}

	//4、获取关联颜色 GoodsColor（带缓存）
	goodsColor := []model.GoodsColor{}
	goods.GoodsColor = strings.ReplaceAll(goods.GoodsColor, "，", ",")
	colorIds := strings.Split(goods.GoodsColor, ",")
	colorCacheKey := fmt.Sprintf("goods_color_%d", id)
	if hasColor := util.CacheDb.Get(colorCacheKey, &goodsColor); !hasColor {
		if len(colorIds) > 0 && colorIds[0] != "" {
			util.DB.Where("id in ?", colorIds).Find(&goodsColor)
		}
		util.CacheDb.Set(colorCacheKey, goodsColor, 3600)
	}

	//5、获取关联配件 GoodsFitting（带缓存）
	goodsFitting := []model.Goods{}
	goods.GoodsFitting = strings.ReplaceAll(goods.GoodsFitting, "，", ",")
	fittingIds := strings.Split(goods.GoodsFitting, ",")
	fittingCacheKey := fmt.Sprintf("goods_fitting_%d", id)
	if hasFitting := util.CacheDb.Get(fittingCacheKey, &goodsFitting); !hasFitting {
		if len(fittingIds) > 0 && fittingIds[0] != "" {
			util.DB.Where("id in ?", fittingIds).Select("id,title,price,goods_version").Find(&goodsFitting)
		}
		util.CacheDb.Set(fittingCacheKey, goodsFitting, 3600)
	}

	//6、获取商品关联的图片 GoodsImage（带缓存）
	goodsImage := []model.GoodsImage{}
	imageCacheKey := fmt.Sprintf("goods_image_%d", id)
	if hasImage := util.CacheDb.Get(imageCacheKey, &goodsImage); !hasImage {
		util.DB.Where("goods_id=?", goods.Id).Limit(6).Find(&goodsImage)
		util.CacheDb.Set(imageCacheKey, goodsImage, 3600)
	}

	//7、获取规格参数信息 GoodsAttr（带缓存）
	goodsAttr := []model.GoodsAttr{}
	attrCacheKey := fmt.Sprintf("goods_attr_%d", id)
	if hasAttr := util.CacheDb.Get(attrCacheKey, &goodsAttr); !hasAttr {
		util.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)
		util.CacheDb.Set(attrCacheKey, goodsAttr, 3600)
	}

	//8、获取更多属性
	// goodsAttrStr := "尺寸:41,42,43|套餐:套餐1,套餐2"

	goodsAttrStr := goods.GoodsAttr
	goodsAttrStr = strings.ReplaceAll(goodsAttrStr, "，", ",")
	goodsAttrStr = strings.ReplaceAll(goodsAttrStr, "：", ":")

	var goodsItemAttrList []model.GoodsItemAttr
	if strings.Contains(goodsAttrStr, ":") {
		goodsAttrStrSlice := strings.Split(goodsAttrStr, "|")
		//创建切片的存储空间
		goodsItemAttrList = make([]model.GoodsItemAttr, len(goodsAttrStrSlice))
		for i := 0; i < len(goodsAttrStrSlice); i++ {
			// 分割冒号左右 左侧为分类 右侧为可选属性
			tempSlice := strings.Split(goodsAttrStrSlice[i], ":")
			goodsItemAttrList[i].Cate = tempSlice[0]
			listSlice := strings.Split(tempSlice[1], ",")
			goodsItemAttrList[i].List = listSlice
		}
	}

	tpl := "shopWeb/product/detail.html"

	con.Render(c, tpl, gin.H{
		"goods":             goods,
		"relationGoods":     relationGoods,
		"goodsGift":         goodsGift,
		"goodsColor":        goodsColor,
		"goodsFitting":      goodsFitting,
		"goodsImage":        goodsImage,
		"goodsAttr":         goodsAttr,
		"goodsItemAttrList": goodsItemAttrList,
	})
}

func (con ProductController) GetImgList(c *gin.Context) {

	goodsId, err1 := util.Int(c.Query("goods_id"))
	colorId, err2 := util.Int(c.Query("color_id"))

	//查询商品图库信息
	goodsImageList := []model.GoodsImage{}
	err3 := util.DB.Where("goods_id=? AND color_id=?", goodsId, colorId).Find(&goodsImageList).Error

	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(200, gin.H{
			"success": false,
			"result":  "",
			"message": "参数错误",
		})
	} else {

		//判断 goodsImageList的长度 如果goodsImageList没有数据，那么我们需要返回当前商品所有的图库信息
		if len(goodsImageList) == 0 {
			util.DB.Where("goods_id=?", goodsId).Find(&goodsImageList)
		}
		c.JSON(200, gin.H{
			"success": true,
			"result":  goodsImageList,
			"message": "获取数据成功",
		})
	}
}
