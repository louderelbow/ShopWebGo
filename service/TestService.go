package service

import (
	"ShopWebGo/config"
	"ShopWebGo/model"
)

type TestData interface{}

// GetTestData 业务逻辑方法（被控制器调用）
func GetTestData() TestData {
	data := model.Admin{}
	// 手动指定表名 + 查询第一条数据
	config.DB.Table("admin").First(&data)
	return data
}
