package middleware

import (
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"
)

func Recover(r *ghttp.Request) {
    defer func() {
        if err := recover(); err != nil {
            // 打日志
            g.Log().Errorf(r.Context(), "panic recovered: %v", err)
            // 统一返回
            r.Response.WriteJson(g.Map{
                "code": 500,
                "msg":  "服务器内部错误，请联系管理员",
                "data": nil,
            })
            // 防止后续 handler 再写响应
            r.Exit()
        }
    }()
    r.Middleware.Next()
}
