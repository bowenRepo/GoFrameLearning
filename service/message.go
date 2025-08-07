package service

import (
	"LeaveWordDemo/model"
	"context"
)

func AddMessage(ctx context.Context, name, content string) (*model.Message, error) {
	// 简单业务规则校验
	if name == "" || content == "" {
		return nil, ErrInvalidParam
	}
	return model.AddMessage(ctx, name, content)
}

func GetMessageList(ctx context.Context, pageNum, pageSize int) ([]*model.Message, int, error) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 || pageSize > 50 {
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	return model.ListMessages(ctx, offset, pageSize)
}

// DeleteMessage 封装删除逻辑
func DeleteMessage(ctx context.Context, id int) error {
	if id <= 0 {
		return ErrInvalidParam
	}
	ok, _ := model.DeleteMessageByID(ctx, id)
	if !ok {
		return ErrMessageNotFound
	}
	return nil
}

// UpdateMessage 校验并调用 model 更新
func UpdateMessage(ctx context.Context, id int, name, content string) (*model.Message, error) {
	if id <= 0 || name == "" || content == "" {
		return nil, ErrInvalidParam
	}
	updated, _ := model.UpdateMessageByID(ctx, id, name, content)
	if updated == nil {
		return nil, ErrMessageNotFound
	}
	return updated, nil
}

// GetMessage 查询单条留言
func GetMessage(ctx context.Context, id int) (*model.Message, error) {
	if id <= 0 {
		return nil, ErrInvalidParam
	}
	msg, _ := model.GetMessageByID(ctx, id)
	if msg == nil {
		return nil, ErrMessageNotFound
	}
	return msg, nil
}
