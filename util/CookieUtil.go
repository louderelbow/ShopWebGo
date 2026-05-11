package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// 定义结构体  缓存结构体 私有
// 定义结构体  缓存结构体 私有
type ginCookie struct{}

// 写入数据的方法
func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {
	//读取.ini里面的数据库配置
	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	des_key := config.Section("").Key("des_key").String()

	bytes, _ := json.Marshal(value)
	//des加密
	desKey := []byte(des_key) //注意：key必须是8位
	encData, _ := DesEncrypt(bytes, desKey)
	c.SetCookie(key, string(encData), 3600*24*30, "/", c.Request.Host, false, true)
}

// 获取数据的方法
func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {

	//读取.ini里面的数据库配置
	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	des_key := config.Section("").Key("des_key").String()

	valueStr, err1 := c.Cookie(key)
	if err1 == nil && valueStr != "" && valueStr != "[]" {
		//des解密
		desKey := []byte(des_key) //注意：key必须是8位
		decData, e := DesDecrypt([]byte(valueStr), desKey)
		if e != nil {
			return false
		} else {
			err2 := json.Unmarshal([]byte(decData), obj)
			return err2 == nil
		}

	}
	return false
}
func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	c.SetCookie(key, "", -1, "/", c.Request.Host, false, true)
	return true
}

// 实例化结构体
var Cookie = &ginCookie{}
