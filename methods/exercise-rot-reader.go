package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return (b-'A'+13)%26 + 'A'
	case 'a' <= b && b <= 'z':
		return (b-'a'+13)%26 + 'a'
	default:
		return b
	}
}

func (r rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
