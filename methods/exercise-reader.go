package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main18() {
	reader.Validate(MyReader{})
}
