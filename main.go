package main

import (
	"ShopWebGo/router"
	"ShopWebGo/util"
	"ShopWebGo/util/middlewares"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {

	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	sessionKey := config.Section("").Key("session_key").String()

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"UnixToTime": util.UnixToTime,
		"Str2Html":   util.Str2Html,
		"FormatImg":  util.FormatImg,
		"Sub":        util.Sub,
		"Mul":        util.Mul,
		"Substr":     util.Substr,
		"FormatAttr": util.FormatAttr,
	})

	r.LoadHTMLGlob("templates/**/**/*")
	r.Static("/static", "./static")

	store := cookie.NewStore([]byte(sessionKey))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middlewares.Cors())

	router.AdminRoutersInit(r)
	router.DefaultRoutersInit(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		fmt.Println("服务启动在 http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务启动失败: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("正在关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务强制关闭:", err)
	}

	if sqlDB, err := util.DB.DB(); err == nil {
		sqlDB.Close()
		fmt.Println("数据库连接已关闭")
	}

	if util.RedisDb != nil {
		util.RedisDb.Close()
		fmt.Println("Redis连接已关闭")
	}

	fmt.Println("服务已安全关闭")
}
