package handler

import (
	"LeaveWordDemo/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func GetMessageList(r *ghttp.Request) {
	pageNum := r.Get("pageNum").Int()
	pageSize := r.Get("pageSize").Int()

	g.Log().Infof(r.Context(), "pageNum=%d, pageSize=%d", pageNum, pageSize)

	list, total, err := service.GetMessageList(r.Context(), pageNum, pageSize)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "msg": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{
		"code": 0,
		"msg":  "ok",
		"data": g.Map{
			"list":  list,
			"total": total,
		},
	})
}

func AddMessage(r *ghttp.Request) {
	var req struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	if err := r.Parse(&req); err != nil {
		r.Response.WriteJson(g.Map{"code": 1, "msg": "参数解析失败"})
		return
	}
	msg, err := service.AddMessage(r.Context(), req.Name, req.Content)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 2, "msg": err.Error()})
		return
	}
	r.Response.WriteJson(g.Map{
		"code": 0,
		"msg":  "留言成功",
		"data": msg,
	})
}

func TestPanic(r *ghttp.Request) {
	panic("I am a test panic!")
}

func DeleteMessage(r *ghttp.Request) {
	id := r.Get("id").Int() // 路径参数
	if err := service.DeleteMessage(r.Context(), id); err != nil {
		code := 1
		msg := err.Error()
		// 自定义不同错误码
		if err == service.ErrMessageNotFound {
			code = 2
		}
		r.Response.WriteJson(g.Map{"code": code, "msg": msg})
		return
	}
	r.Response.WriteJson(g.Map{"code": 0, "msg": "删除成功"})
}

func UpdateMessage(r *ghttp.Request) {
    id := r.Get("id").Int()
    var req struct {
        Name    string `json:"name"`
        Content string `json:"content"`
    }
    if err := r.Parse(&req); err != nil {
        r.Response.WriteJson(g.Map{"code": 1, "msg": "参数解析失败"})
        return
    }
    msg, err := service.UpdateMessage(r.Context(), id, req.Name, req.Content)
    if err != nil {
        code := 2
        if err == service.ErrInvalidParam {
            code = 3
        }
        if err == service.ErrMessageNotFound {
            code = 4
        }
        r.Response.WriteJson(g.Map{"code": code, "msg": err.Error()})
        return
    }
    r.Response.WriteJson(g.Map{"code": 0, "msg": "更新成功", "data": msg})
}


func GetMessage(r *ghttp.Request) {
    id := r.Get("id").Int()  // 路径参数名要和 router 保持一致
    msg, err := service.GetMessage(r.Context(), id)
    if err != nil {
        code := 1
        if err == service.ErrInvalidParam {
            code = 2
        }
        if err == service.ErrMessageNotFound {
            code = 3
        }
        r.Response.WriteJson(g.Map{"code": code, "msg": err.Error()})
        return
    }
    r.Response.WriteJson(g.Map{"code": 0, "msg": "ok", "data": msg})
}
