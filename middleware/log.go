package middleware

import (
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"
)

// 每当用户发起请求，就在后台日志中记录一行信息
func Logger(r *ghttp.Request) {
    g.Log().Infof(r.Context(), "[%s] %s %s", r.Method, r.URL.Path, r.GetBodyString())
    r.Middleware.Next()
}
