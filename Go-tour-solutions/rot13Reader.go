package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// 将得到的字母用rot13方法转换
func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b += 13
	case 'a' <= b && b <= 'm':
		b += 13
	case 'N' <= b && b <= 'Z':
		b -= 13
	case 'n' <= b && b <= 'z':
		b -= 13
	}
	return b
}

// 重写Read方法
func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := range p {
		p[i] = rot13(p[i])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
