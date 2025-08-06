package model

type Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// 先用内存模拟留言列表
var messages = []*Message{}
var nextID = 1

// 添加留言
func AddMessage(name, content string) *Message {
	msg := &Message{
		ID:      nextID,
		Name:    name,
		Content: content,
	}
	messages = append(messages, msg)
	nextID++
	return msg
}

// 查询留言（分页）: 指针 + 大小
func ListMessages(offset, limit int) ([]*Message, int) {
	total := len(messages)
	if offset > total {
		return []*Message{}, total
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return messages[offset:end], total
}


// 删除指定 ID 的留言，成功返回 true，否则 false
func DeleteMessageByID(id int) bool {
    for i, m := range messages {
        if m.ID == id {
            // 删掉这个元素
            messages = append(messages[:i], messages[i+1:]...)
            return true
        }
    }
    return false
}

// 更新指定 ID 的留言内容，成功返回更新后的 *Message，否则 nil
func UpdateMessageByID(id int, name, content string) *Message {
    for _, m := range messages {
        if m.ID == id {
            m.Name    = name
            m.Content = content
            return m
        }
    }
    return nil
}
