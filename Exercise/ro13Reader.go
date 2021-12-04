package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	// A:65 to Z:90
	// a:97 to z:122
	bt, err := r.r.Read(b)

	for i := range b {
		if (b[i] >= 65 && b[i] <= 77) || (b[i] >= 97 && b[i] <= 109) {
			b[i] = b[i] + 13
		} else if (b[i] >= 78 && b[i] <= 90) || (b[i] >= 110 && b[i] <= 122) {
			b[i] = b[i] - 13
		}
	}
	return bt, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
