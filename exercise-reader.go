package main

import (
	"errors"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(b []byte) (int, error) {
	l := len(b)
	if l == 0 {
		return 0, errors.New("no space to read")
	}
	for i := 0; i < l; i++ {
		b[i] = 65
	}
	return l, nil
}

func main() {
	reader.Validate(MyReader{})
}
