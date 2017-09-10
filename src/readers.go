package main

import (
	"fmt"
	"io"
	"strings"
)

func test1() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		fmt.Printf("b[:%d] = %q\n", n, b[:n])
		if err == io.EOF {
			fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
			break
		}
	}
}

/*
	Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/
type MyReader struct{
	size int // size control, so it's finite
}

func (m *MyReader) Read(b []byte) (i int, e error) {
    for x := range b {
        b[x] = 'A'
    }
    return len(b), nil
}

func (m *MyReader) Validate() (ok bool) {
	if m.size < 0 {
    	return false
	}
	m.size--
    return true
}

func main() {
	reader := MyReader{100}
	b := make([]byte, 1) 
    for reader.Validate() == true {
		n, err := reader.Read(b)
		fmt.Printf("b[:%d] = %q\n", n, b[:n])
		if err == io.EOF {
			fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
			break
		}
	}
}
