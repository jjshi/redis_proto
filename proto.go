package proto

import (
	"strconv"
)

// 定义响应协议
const (
	ErrorReply  = '-'
	StatusReply = '+'
	IntReply    = ':'
	StringReply = '$'
	ArrayReply  = '*'
)

// RedisError 错误协议结构
type RedisError string

func (e RedisError) Error() string {
	return string(e)
}

// 解析错误
func parseError(data []byte) error {
	return RedisError(string(data[1:]))
}

// 解析整形
func parseInt(s string) (int, error) {
	r, err := strconv.ParseInt(s, 10, 64)
	return int(r), err
}
