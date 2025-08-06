package middleware

import (
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/gogf/gf/v2/frame/g"
)

const demoToken = "abcdef123456" // 测试环境用

func Auth(r *ghttp.Request) {
    token := r.Header.Get("Authorization")
    if token != demoToken {
        r.Response.WriteJson(g.Map{
            "code": 401,
            "msg":  "未授权，缺少或无效Token",
            "data": nil,
        })
        r.ExitAll()
        return
    }
    r.Middleware.Next()
}
