package util

import (
	"ShopWebGo/model"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"reflect"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// Oss上传
func OssUpload(file *multipart.FileHeader, dst string) (string, error) {

	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	//读取.ini里面的数据库配置
	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	accessKeyID := config.Section("oss").Key("access_key_id").String()
	accessKeySecret := config.Section("oss").Key("access_key_secret").String()

	// 创建OSSClient实例。
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", accessKeyID, accessKeySecret)
	if err != nil {
		return "", err
	}

	// 获取存储空间。
	bucket, err := client.Bucket("java-system")
	if err != nil {
		return "", err
	}

	// 上传文件流。
	err = bucket.PutObject(dst, f)
	if err != nil {
		return "", err
	}
	return dst, nil
}

// 获取Oss的状态
func GetOssStatus() int {
	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	ossStatus, _ := Int(config.Section("oss").Key("status").String())
	return ossStatus
}

// 上传图片
func UploadImg(c *gin.Context, picName string) (string, error) {
	ossStatus := GetOssStatus()
	if ossStatus == 1 {
		return OssUploadImg(c, picName)
	} else {
		return LocalUploadImg(c, picName)
	}

}

// 上传图片
func LocalUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、创建图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "./static/upload/" + day

	err1 := os.MkdirAll(dir, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil

}

// 上传图片到Oss
func OssUploadImg(c *gin.Context, picName string) (string, error) {
	// 1、获取上传的文件
	file, err := c.FormFile(picName)

	if err != nil {
		return "", err
	}

	// 2、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3、定义图片保存目录  static/upload/20210624

	day := GetDay()
	dir := "static/upload/" + day

	// 4、生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnixNano(), 10) + extName

	// 5、执行上传
	dst := path.Join(dir, fileName)

	_, err = OssUpload(file, dst)
	if err != nil {
		return "", err
	}
	return dst, nil

}

// 通过列获取值
func GetSettingFromColumn(columnName string) string {
	//redis file
	setting := model.Setting{}
	DB.First(&setting)
	//反射来获取
	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}

// 格式化输出图片
func FormatImg(str string) string {
	ossStatus := GetOssStatus()
	if ossStatus == 1 {
		return GetSettingFromColumn("OssDomain") + str
	} else {
		return "/" + str
	}
}
