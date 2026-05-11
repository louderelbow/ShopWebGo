package admin

import (
	"ShopWebGo/model"
	"ShopWebGo/util"
	"fmt"
	"math"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	//当前页数
	page, _ := util.Int(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//每页查询的数量
	pageSize := 5

	//获取keyword
	keyword := c.Query("keyword")

	goodsList := []model.Goods{}
	listDB := util.DB.Where("is_delete = 0")
	if len(keyword) > 0 {
		listDB = listDB.Where("title LIKE ?", "%"+keyword+"%")
	}
	listDB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsList)

	var count int64
	countDB := util.DB.Table("goods").Where("is_delete = 0")
	if len(keyword) > 0 {
		countDB = countDB.Where("title LIKE ?", "%"+keyword+"%")
	}
	countDB.Count(&count)

	//判断最后一页有没有数据 如果没有跳转到第一页
	if len(goodsList) > 0 {
		c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
			"goodsList": goodsList,
			//注意float64类型
			"totalPages": math.Ceil(float64(count) / float64(pageSize)),
			"page":       page,
			"keyword":    keyword,
		})
	} else {
		if page != 1 {
			c.Redirect(302, "/admin/goods")
		} else {
			c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
				"goodsList": goodsList,
				//注意float64类型
				"totalPages": math.Ceil(float64(count) / float64(pageSize)),
				"page":       page,
				"keyword":    keyword,
			})
		}

	}

}

func (con GoodsController) Add(c *gin.Context) {
	//获取商品分类
	goodsCateList := []model.GoodsCate{}
	util.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	//获取所有颜色信息
	goodsColorList := []model.GoodsColor{}
	util.DB.Find(&goodsColorList)

	//获取商品规格包装
	goodsTypeList := []model.GoodsType{}
	util.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
	})
}

func (con GoodsController) GoodsTypeAttribute(c *gin.Context) {
	cateId, err1 := util.Int(c.Query("cateId"))
	goodsTypeAttributeList := []model.GoodsTypeAttribute{}
	err2 := util.DB.Where("cate_id = ?", cateId).Find(&goodsTypeAttributeList).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  goodsTypeAttributeList,
		})
	}
}

func (con GoodsController) DoAdd(c *gin.Context) {

	//1、获取表单提交过来的数据 进行判断

	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := util.Int(c.PostForm("cate_id"))
	goodsNumber, _ := util.Int(c.PostForm("goods_number"))
	//注意小数点
	marketPrice, _ := util.Float(c.PostForm("market_price"))
	price, _ := util.Float(c.PostForm("price"))

	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	//获取的是切片
	goodsColorArr := c.PostFormArray("goods_color")

	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := util.Int(c.PostForm("is_delete"))
	isHot, _ := util.Int(c.PostForm("is_hot"))
	isBest, _ := util.Int(c.PostForm("is_best"))
	isNew, _ := util.Int(c.PostForm("is_new"))
	goodsTypeId, _ := util.Int(c.PostForm("goods_type_id"))
	sort, _ := util.Int(c.PostForm("sort"))
	status, _ := util.Int(c.PostForm("status"))
	addTime := int(util.GetUnix())

	//2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")

	//3、上传图片   生成缩略图
	goodsImg, _ := util.UploadImg(c, "goods_img")
	if len(goodsImg) > 0 {
		//判断 本地图片才需要处理
		if util.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				util.ResizeGoodsImage(goodsImg)
				wg.Done()
			}()
		}

	}
	//4、增加商品数据

	goods := model.Goods{
		Title:         title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    100,
		GoodsNumber:   goodsNumber,
		MarketPrice:   marketPrice,
		Price:         price,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsKeywords: goodsKeywords,
		GoodsDesc:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       addTime,
		GoodsColor:    goodsColorStr,
		GoodsImg:      goodsImg,
	}
	err := util.DB.Create(&goods).Error
	if err != nil {
		con.error(c, "增加失败", "/admin/goods/add")
	}
	//5、增加图库 信息
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := model.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(util.GetUnix())
			util.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()
	//6、增加规格包装
	wg.Add(1)
	go func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, attributeIdErr := util.Int(attrIdList[i])
			if attributeIdErr == nil {
				//获取商品类型属性的数据
				goodsTypeAttributeObj := model.GoodsTypeAttribute{Id: goodsTypeAttributeId}
				util.DB.Find(&goodsTypeAttributeObj)
				//给商品属性里面增加数据  规格包装
				goodsAttrObj := model.GoodsAttr{}
				goodsAttrObj.GoodsId = goods.Id
				goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
				goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
				goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
				goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
				goodsAttrObj.AttributeValue = attrValueList[i]
				goodsAttrObj.Status = 1
				goodsAttrObj.Sort = 10
				goodsAttrObj.AddTime = int(util.GetUnix())
				util.DB.Create(&goodsAttrObj)
			}

		}
		wg.Done()
	}()
	wg.Wait()
	con.success(c, "增加数据成功", "/admin/goods")
}

// 修改
func (con GoodsController) Edit(c *gin.Context) {

	// 1、获取要修改的商品数据
	id, err := util.Int(c.Query("id"))
	if err != nil {
		con.error(c, "传入参数错误", "/admin/goods")
	}
	goods := model.Goods{Id: id}
	util.DB.Find(&goods)

	// 2、获取商品分类
	goodsCateList := []model.GoodsCate{}
	util.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	// 3、获取所有颜色 以及选中的颜色
	goodsColorSlice := strings.Split(goods.GoodsColor, ",")
	goodsColorMap := make(map[string]string)
	for _, v := range goodsColorSlice {
		goodsColorMap[v] = v
	}

	goodsColorList := []model.GoodsColor{}
	util.DB.Find(&goodsColorList)
	for i := 0; i < len(goodsColorList); i++ {
		if _, ok := goodsColorMap[util.String(goodsColorList[i].Id)]; ok {
			goodsColorList[i].Checked = true
		}
	}

	// 4、商品的图库信息
	goodsImageList := []model.GoodsImage{}
	util.DB.Where("goods_id=?", goods.Id).Find(&goodsImageList)

	// 5、获取商品类型
	goodsTypeList := []model.GoodsType{}
	util.DB.Find(&goodsTypeList)

	// 6、获取规格信息
	goodsAttr := []model.GoodsAttr{}
	util.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)
	goodsAttrStr := ""

	for _, v := range goodsAttr {
		if v.AttributeType == 1 {
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: </span> <input type="hidden" name="attr_id_list" value="%v" />   <input type="text" name="attr_value_list" value="%v" /></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else if v.AttributeType == 2 {
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span><input type="hidden" name="attr_id_list" value="%v" />  <textarea cols="50" rows="3" name="attr_value_list">%v</textarea></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else {
			//获取当前类型对应的值
			goodsTypeArttribute := model.GoodsTypeAttribute{Id: v.AttributeId}
			util.DB.Find(&goodsTypeArttribute)
			attrValueSlice := strings.Split(goodsTypeArttribute.AttrValue, "\n")

			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" /> `, v.AttributeTitle, v.AttributeId)
			goodsAttrStr += fmt.Sprintf(`<select name="attr_value_list">`)
			for i := 0; i < len(attrValueSlice); i++ {
				if attrValueSlice[i] == v.AttributeValue {
					goodsAttrStr += fmt.Sprintf(`<option value="%v" selected >%v</option>`, attrValueSlice[i], attrValueSlice[i])
				} else {
					goodsAttrStr += fmt.Sprintf(`<option value="%v">%v</option>`, attrValueSlice[i], attrValueSlice[i])
				}
			}
			goodsAttrStr += fmt.Sprintf(`</select>`)
			goodsAttrStr += fmt.Sprintf(`</li>`)

		}
	}

	//获取上一页的地址
	fmt.Println(c.Request.Referer())

	c.HTML(http.StatusOK, "admin/goods/edit.html", gin.H{
		"goods":          goods,
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
		"goodsAttrStr":   goodsAttrStr,
		"goodsImageList": goodsImageList,
		"prevPage":       c.Request.Referer(), //获取上一页的地址
	})
}
func (con GoodsController) DoEdit(c *gin.Context) {

	//1、获取表单提交过来的数据
	id, err1 := util.Int(c.PostForm("id"))
	if err1 != nil {
		con.error(c, "传入参数错误", "/admin/goods")
	}
	//获取上一页的地址
	prevPage := c.PostForm("prevPage")
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := util.Int(c.PostForm("cate_id"))
	goodsNumber, _ := util.Int(c.PostForm("goods_number"))
	//注意小数点
	marketPrice, _ := util.Float(c.PostForm("market_price"))
	price, _ := util.Float(c.PostForm("price"))
	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	//获取的是切片
	goodsColorArr := c.PostFormArray("goods_color")
	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := util.Int(c.PostForm("is_delete"))
	isHot, _ := util.Int(c.PostForm("is_hot"))
	isBest, _ := util.Int(c.PostForm("is_best"))
	isNew, _ := util.Int(c.PostForm("is_new"))
	goodsTypeId, _ := util.Int(c.PostForm("goods_type_id"))
	sort, _ := util.Int(c.PostForm("sort"))
	status, _ := util.Int(c.PostForm("status"))

	//2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")
	//3、修改数据
	goods := model.Goods{Id: id}
	util.DB.Find(&goods)
	goods.Title = title
	goods.SubTitle = subTitle
	goods.GoodsSn = goodsSn
	goods.CateId = cateId
	goods.GoodsNumber = goodsNumber
	goods.MarketPrice = marketPrice
	goods.Price = price
	goods.RelationGoods = relationGoods
	goods.GoodsAttr = goodsAttr
	goods.GoodsVersion = goodsVersion
	goods.GoodsGift = goodsGift
	goods.GoodsFitting = goodsFitting
	goods.GoodsKeywords = goodsKeywords
	goods.GoodsDesc = goodsDesc
	goods.GoodsContent = goodsContent
	goods.IsDelete = isDelete
	goods.IsHot = isHot
	goods.IsBest = isBest
	goods.IsNew = isNew
	goods.GoodsTypeId = goodsTypeId
	goods.Sort = sort
	goods.Status = status
	goods.GoodsColor = goodsColorStr

	//4、上传图片   生成缩略图
	goodsImg, err2 := util.UploadImg(c, "goods_img")
	if err2 == nil && len(goodsImg) > 0 {
		goods.GoodsImg = goodsImg
		if util.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				util.ResizeGoodsImage(goodsImg)
				wg.Done()
			}()
		}
	}

	err3 := util.DB.Save(&goods).Error
	if err3 != nil {
		con.error(c, "修改失败", "/admin/goods/edit?id="+util.String(id))
		return
	}

	//5、修改图库 增加图库信息
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := model.GoodsImage{}
			goodsImgObj.GoodsId = goods.Id
			goodsImgObj.ImgUrl = v
			goodsImgObj.Sort = 10
			goodsImgObj.Status = 1
			goodsImgObj.AddTime = int(util.GetUnix())
			util.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()
	//6、修改规格包装  1、删除当前商品下面的规格包装   2、重新执行增加

	// 6.1删除当前商品下面的规格包装
	goodsAttrObj := model.GoodsAttr{}
	util.DB.Where("goods_id=?", goods.Id).Delete(&goodsAttrObj)
	//6.2、重新执行增加
	wg.Add(1)
	go func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, attributeIdErr := util.Int(attrIdList[i])
			if attributeIdErr == nil {
				//获取商品类型属性的数据
				goodsTypeAttributeObj := model.GoodsTypeAttribute{Id: goodsTypeAttributeId}
				util.DB.Find(&goodsTypeAttributeObj)
				//给商品属性里面增加数据  规格包装
				goodsAttrObj := model.GoodsAttr{}
				goodsAttrObj.GoodsId = goods.Id
				goodsAttrObj.AttributeTitle = goodsTypeAttributeObj.Title
				goodsAttrObj.AttributeType = goodsTypeAttributeObj.AttrType
				goodsAttrObj.AttributeId = goodsTypeAttributeObj.Id
				goodsAttrObj.AttributeCateId = goodsTypeAttributeObj.CateId
				goodsAttrObj.AttributeValue = attrValueList[i]
				goodsAttrObj.Status = 1
				goodsAttrObj.Sort = 10
				goodsAttrObj.AddTime = int(util.GetUnix())
				util.DB.Create(&goodsAttrObj)
			}

		}
		wg.Done()
	}()
	wg.Wait()
	if len(prevPage) > 0 {
		con.success(c, "修改数据成功", prevPage)
	} else {
		con.success(c, "修改数据成功", "/admin/goods")
	}

}

// 富文本编辑器上传图片
func (con GoodsController) EditorImageUpload(c *gin.Context) {
	//上传图片
	imgDir, err := util.UploadImg(c, "file") //注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		if util.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				util.ResizeGoodsImage(imgDir)
				wg.Done()
			}()
			c.JSON(http.StatusOK, gin.H{
				"link": imgDir,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"link": util.GetSettingFromColumn("OssDomain") + imgDir,
			})
		}

	}
}

func (con GoodsController) GoodsImageUpload(c *gin.Context) {
	imgDir, err := util.UploadImg(c, "file") //注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		if util.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				util.ResizeGoodsImage(imgDir)
				wg.Done()
			}()
		}

		c.JSON(http.StatusOK, gin.H{
			"link": imgDir,
		})
	}
}

// 修改商品图库关联的颜色
func (con GoodsController) ChangeGoodsImageColor(c *gin.Context) {
	//获取图片id 获取颜色id
	goodsImageId, err1 := util.Int(c.Query("goods_image_id"))
	colorId, err2 := util.Int(c.Query("color_id"))
	goodsImage := model.GoodsImage{Id: goodsImageId}
	util.DB.Find(&goodsImage)
	goodsImage.ColorId = colorId
	err3 := util.DB.Save(&goodsImage).Error
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "更新失败",
			"success": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":  "更新成功",
			"success": true,
		})
	}

}

// 删除图库
func (con GoodsController) RemoveGoodsImage(c *gin.Context) {
	//获取图片id
	goodsImageId, err1 := util.Int(c.Query("goods_image_id"))
	goodsImage := model.GoodsImage{Id: goodsImageId}
	err2 := util.DB.Delete(&goodsImage).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "删除失败",
			"success": false,
		})
	} else {
		//删除图片
		// os.Remove()
		c.JSON(http.StatusOK, gin.H{
			"result":  "删除成功",
			"success": true,
		})
	}

}

// 删除数据
func (con GoodsController) Delete(c *gin.Context) {
	id, err := util.Int(c.Query("id"))
	if err != nil {
		con.error(c, "传入数据错误", "/admin/goods")
	} else {
		goods := model.Goods{Id: id}
		util.DB.Find(&goods)
		goods.IsDelete = 1
		goods.Status = 0
		util.DB.Save(&goods)
		//获取上一页
		prevPage := c.Request.Referer()
		if len(prevPage) > 0 {
			con.success(c, "删除数据成功", prevPage)
		} else {
			con.success(c, "删除数据成功", "/admin/goods")
		}

	}

}
