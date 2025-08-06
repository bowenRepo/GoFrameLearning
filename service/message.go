package service

import "LeaveWordDemo/model"

func AddMessage(name, content string) (*model.Message, error) {
    // 简单业务规则校验
    if name == "" || content == "" {
        return nil, ErrInvalidParam
    }
    return model.AddMessage(name, content), nil
}

func GetMessageList(pageNum, pageSize int) ([]*model.Message, int, error) {
    if pageNum <= 0 {
        pageNum = 1
    }
    if pageSize <= 0 || pageSize > 50 {
        pageSize = 10
    }
    offset := (pageNum - 1) * pageSize
    msgs, total := model.ListMessages(offset, pageSize)
    return msgs, total, nil
}

// DeleteMessage 封装删除逻辑
func DeleteMessage(id int) error {
    if id <= 0 {
        return ErrInvalidParam
    }
    ok := model.DeleteMessageByID(id)
    if !ok {
        return ErrMessageNotFound
    }
    return nil
}

// UpdateMessage 校验并调用 model 更新
func UpdateMessage(id int, name, content string) (*model.Message, error) {
    if id <= 0 || name == "" || content == "" {
        return nil, ErrInvalidParam
    }
    updated := model.UpdateMessageByID(id, name, content)
    if updated == nil {
        return nil, ErrMessageNotFound
    }
    return updated, nil
}


// GetMessage 查询单条留言
func GetMessage(id int) (*model.Message, error) {
    if id <= 0 {
        return nil, ErrInvalidParam
    }
    msg := model.GetMessageByID(id)
    if msg == nil {
        return nil, ErrMessageNotFound
    }
    return msg, nil
}
