package proto

import (
	"bufio"
	"errors"
	"io"
)

// Reader 协议结构
type Reader struct {
	rd *bufio.Reader
}

// NewReader 实例化reader
func NewReader(rd io.Reader) *Reader {
	return &Reader{
		rd: bufio.NewReader(rd),
	}
}

// 读取一行数据
func (r *Reader) readLine() ([]byte, error) {
	line, isPrefix, err := r.rd.ReadLine()
	if err != nil {
		return nil, err
	}

	if isPrefix {
		return nil, bufio.ErrBufferFull
	}

	if len(line) == 0 {
		return nil, errors.New("data empty")
	}

	return line, nil
}

// Parse 解析
func (r *Reader) Parse() (interface{}, error) {
	data, err := r.readLine()
	if err != nil {
		return nil, err
	}

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
