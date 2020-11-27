package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (pr *rot13Reader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, errors.New("no space to read")
	}
	n, err := pr.r.Read(b)
	for i := 0; i < n; i++ {
		bi := b[i]
		var base byte
		if 65 <= bi && bi < 65+26 {
			base = 65
		} else if 97 <= bi && bi < 97+26 {
			base = 97
		}
		if base > 0 {
			offset := bi - base
			b[i] = base + (offset+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, _ = io.Copy(os.Stdout, &r)
}
