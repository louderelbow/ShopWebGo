package util

import (
	"fmt"
	"os"
	"path"
	"strings"

	. "github.com/hunterhug/go_image"
)

// 生成商品缩略图
func ResizeGoodsImage(filename string) {
	// 1. 判断文件是否存在（防止本地无文件）
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("缩略图生成失败：文件不存在 ->", filename)
		return
	}
	ThumbnailSize := strings.ReplaceAll(GetSettingFromColumn("ThumbnailSize"), "，", ",")
	thumbnailSizeSlice := strings.Split(ThumbnailSize, ",")
	extname := path.Ext(filename)

	// 3. 循环生成缩略图
	for i := 0; i < len(thumbnailSizeSlice); i++ {
		size := strings.TrimSpace(thumbnailSizeSlice[i])
		if size == "" {
			continue
		}
		// 修复尺寸转换（用你自己的util.Int）
		w, err := Int(size)
		if err != nil || w <= 0 {
			continue
		}
		// 修复命名：去掉重复后缀
		rawName := strings.TrimSuffix(filename, extname)
		savepath := fmt.Sprintf("%s_%dx%d%s", rawName, w, w, extname)

		// 生成缩略图
		err = ThumbnailF2F(filename, savepath, w, w)
		if err != nil {
			fmt.Println("缩略图生成失败：", err)
		} else {
			fmt.Println("缩略图生成成功：", savepath)
		}
	}
}
