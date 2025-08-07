package model

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// 表名常量，后续写查询更直观
const Table = "message"

// Message 映射 message 表
type Message struct {
	Id        int    `json:"id"        orm:"id,primary"`
	Name      string `json:"name"      orm:"name"`
	Content   string `json:"content"   orm:"content"`
	CreatedAt string `json:"createdAt" orm:"created_at"`
}

// 添加留言（单条）
func AddMessage(ctx context.Context, name, content string) (*Message, error) {
	msg := &Message{Name: name, Content: content}
	// 直接 Insert 结构体，gdb 自动读取 orm 标签
	r, err := g.DB().Model(Table).Insert(ctx, msg)
	if err != nil {
		return nil, err
	}
	id, _ := r.LastInsertId()
	return GetMessageByID(ctx, int(id))
}

// 查询留言（分页）
func ListMessages(ctx context.Context, offset, limit int) ([]*Message, int, error) {
	var list []*Message
	// 查总数
	total, err := g.DB().Model(Table).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	// 查数据
	err = g.DB().
		Model(Table).
		Order("id DESC").
		Limit(offset, limit).
		Scan(ctx, &list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// 查询单条（id）
func GetMessageByID(ctx context.Context, id int) (*Message, error) {
	var msg *Message
	err := g.DB().Model(Table).Where("id", id).Scan(ctx, &msg)
	return msg, err
}

// delete 单条(id)
func DeleteMessageByID(ctx context.Context, id int) (bool, error) {
	r, err := g.DB().Model(Table).Where("id", id).Delete(ctx)
	if err != nil {
		return false, err
	}
	affected, _ := r.RowsAffected()
	return affected > 0, nil
}

// update 单条(id)
func UpdateMessageByID(ctx context.Context, id int, name, content string) (*Message, error) {
	_, err := g.DB().Model(Table).Data(g.Map{
		"name":    name,
		"content": content,
	}).Where("id", id).Update(ctx)
	if err != nil {
		return nil, err
	}
	return GetMessageByID(ctx, id)
}
