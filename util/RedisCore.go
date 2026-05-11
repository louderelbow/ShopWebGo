package util

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

var ctx = context.Background()
var RedisDb *redis.Client
var redisEnable bool

func init() {

	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	ip := config.Section("redis").Key("ip").String()
	port := config.Section("redis").Key("port").String()
	redisEnable, _ = config.Section("redis").Key("redisEnable").Bool()

	if redisEnable {
		//连接redis数据库
		RedisDb = redis.NewClient(&redis.Options{
			Addr:     ip + ":" + port,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		_, err := RedisDb.Ping(ctx).Result()
		if err != nil {
			fmt.Println("redis数据库连接失败")
		} else {
			fmt.Println("redis数据库连接成功...")
		}
	} else {
		fmt.Println("未连接Redis")
	}

}

// 实现数据处理封装解耦
type cacheDb struct{}

func (c cacheDb) Set(key string, value interface{}, expiration int) {
	if redisEnable {
		v, err := json.Marshal(value)
		if err == nil {
			RedisDb.Set(ctx, key, string(v), time.Second*time.Duration(expiration))
		}
	}
}

func (c cacheDb) Get(key string, obj interface{}) bool {
	if redisEnable {
		valueStr, err1 := RedisDb.Get(ctx, key).Result()
		if err1 == nil && valueStr != "" {
			err2 := json.Unmarshal([]byte(valueStr), obj)
			return err2 == nil
		}
		return false
	}
	return false
}

// 清除缓存
func (c cacheDb) FlushAll() {
	if redisEnable {
		RedisDb.FlushAll(ctx)
	}
}

var CacheDb = &cacheDb{}
