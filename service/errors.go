// 自定义错误
package service

import "errors"
var (
	ErrInvalidParam = errors.New("name or content is empty")
	ErrMessageNotFound = errors.New("留言不存在")
)