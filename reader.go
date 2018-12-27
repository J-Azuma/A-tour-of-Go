package main

import "golang.org/x/tour/reader"

type MyReader struct{}

//ToDo: Add a Read([]byte) (int, erors) method to MyReader.

func (r MyReader) Read(rb []byte) (n int, e error) {
	for n, e = 0, nil; n < len(rb); n++ {
		rb[n] = 'A'
	}
	return len(rb), nil
}
func main() {
	reader.Validate(MyReader{})
}
