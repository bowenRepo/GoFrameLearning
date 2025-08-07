package model

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 表名常量，后续写查询更直观
const Table = "message"

// Message 映射 message 表
type Message struct {
	Id        int         `json:"id"        orm:"id,primary"`
	Name      string      `json:"name"      orm:"name"`
	Content   string      `json:"content"   orm:"content"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at"`
}

// 添加留言
func AddMessage(ctx context.Context, name, content string) (*Message, error) {
	msg := &Message{Name: name, Content: content}
	r, err := g.DB().Model(Table).Ctx(ctx).Insert(msg)
	if err != nil {
		return nil, err
	}
	id, _ := r.LastInsertId()
	return GetMessageByID(ctx, int(id))
}

// 查询留言（分页）
func ListMessages(ctx context.Context, offset, limit int) ([]*Message, int, error) {
	var list []*Message
	total, err := g.DB().Model(Table).Ctx(ctx).Count()
	if err != nil {
		return nil, 0, err
	}
	err = g.DB().Model(Table).Ctx(ctx).Order("id DESC").Limit(offset, limit).Scan(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// 查询单条
func GetMessageByID(ctx context.Context, id int) (*Message, error) {
	var msg *Message
	err := g.DB().Model(Table).Ctx(ctx).Where("id", id).Scan(&msg)
	return msg, err
}

// 删除单条
func DeleteMessageByID(ctx context.Context, id int) (bool, error) {
	r, err := g.DB().Model(Table).Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return false, err
	}
	affected, _ := r.RowsAffected()
	return affected > 0, nil
}

// 更新单条
func UpdateMessageByID(ctx context.Context, id int, name, content string) (*Message, error) {
	_, err := g.DB().Model(Table).Ctx(ctx).
		Data(g.Map{"name": name, "content": content}).
		Where("id", id).Update()
	if err != nil {
		return nil, err
	}
	return GetMessageByID(ctx, id)
}
