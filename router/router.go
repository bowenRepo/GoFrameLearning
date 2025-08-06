// package router

// import (
// 	"LeaveWordDemo/handler"

// 	"github.com/gogf/gf/v2/net/ghttp"
// )

// func Register(s *ghttp.Server) {
// 	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
// 		group.GET("/messages", handler.GetMessageList)
// 		group.POST("/messages", handler.AddMessage)
// 		group.GET("/test_panic", handler.TestPanic)
// 	})
// }

// router/router.go
package router

import (
	"LeaveWordDemo/handler"
	"LeaveWordDemo/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(s *ghttp.Server) {
	// /api/v1 组所有路由都需鉴权
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		group.GET("/messages", handler.GetMessageList)
		group.POST("/messages", handler.AddMessage)
		group.GET("/test_panic", handler.TestPanic)
		group.DELETE("/messages/{id}", handler.DeleteMessage)
        group.PUT("/messages/{id}", handler.UpdateMessage)
        
        // 单条详情接口
        group.GET("/messages/{id}", handler.GetMessage)

	})
}

// func Register(s *ghttp.Server) {
//     // 建立 /api/v1 路由组
//     group := s.Group("/api/v1")

//     // 以下示例：只对 POST、DELETE、PUT 接口加鉴权
//     group.GET("/messages", handler.GetMessageList)                               // 不鉴权，公开读取
//     group.POST("/messages", middleware.Auth, handler.AddMessage)                 // 鉴权，发布需登录
//     group.DELETE("/messages/{id}", middleware.Auth, handler.DeleteMessage)       // 鉴权，删除需登录
//     group.PUT("/messages/{id}", middleware.Auth, handler.UpdateMessage)          // 鉴权，更新需登录
// }
