package util

import (
	"sync"

	"gopkg.in/ini.v1"
)

var (
	configOnce sync.Once
	cachedCfg  *ini.File
)

// GetConfig 返回缓存的 app.ini 配置，全局只加载一次
func GetConfig() *ini.File {
	configOnce.Do(func() {
		var err error
		cachedCfg, err = ini.Load("./config/app.ini")
		if err != nil {
			panic("failed to load config: " + err.Error())
		}
	})
	return cachedCfg
}
