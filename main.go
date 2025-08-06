package main

import (
	"LeaveWordDemo/middleware"
	"LeaveWordDemo/router"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	s := g.Server()

	s.Use(middleware.Recover)
	s.Use(middleware.Logger) // 全局中间件
    s.Use(middleware.Auth)    // 统一鉴权

	router.Register(s)       // 注册路由和 handler
	s.Run()
}
