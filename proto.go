package proto

import (
	"bytes"
	"errors"
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

// 协议结尾符
var tail = []byte("\r\n")

// 截取格式
func format(data []byte) (r []byte, err error) {
	parts := bytes.Trim(data, " ")
	if len(parts) < len(tail) {
		err = errors.New("data too short")
		return
	}

	if ok := bytes.Equal(parts[len(data)-2:], tail); !ok {
		err = errors.New("data format fail")
		return
	}
	return parts[0 : len(data)-2], nil
}

// 解析
func decode(data []byte) (interface{}, error) {
	switch data[0] {
	case ErrorReply:
		return nil, parseError(data)
	case IntReply:
		return parseInt(string(data[1:]))
	case StatusReply, StringReply:
		return string(data[1:]), nil
	case ArrayReply:
		return parseInt(string(data[1:]))
	default:
		return nil, errors.New("not support reply")
	}
}

// ParseProto 解析协议
func ParseProto(protocol []byte) (interface{}, error) {
	data, err := format(protocol)
	if err != nil {
		return nil, err
	}

	return decode(data)
}

// RedisError 错误协议结构
type RedisError string

func (e RedisError) Error() string {
	return string(e)
}

// 解析错误
func parseError(line []byte) error {
	return RedisError(string(line[1:]))
}

// 解析整形
func parseInt(s string) (int, error) {
	r, err := strconv.ParseInt(s, 10, 64)
	return int(r), err
}
