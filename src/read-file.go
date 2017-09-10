package main

import (
    "fmt"
    "os"
    "log"
    "io"
    "bufio"
)

func test1() {
	file, err := os.Open("read-file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)

	for {
		n, err := file.Read(data) 
		if err == io.EOF {
			break
		}
		fmt.Printf("read %d bytes: %s\n", n, data[:n])
	}
}

func test2() {
	file, err := os.Open("read-file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	//data := make([]byte, 1024)
	
    w := bufio.NewReader(file)
	n := 0

	for {
		data, is_prefix, err := w.ReadLine()
		if err == io.EOF {
			break
		}
		if is_prefix == true {
			fmt.Println("[got a long line]")
		}
		n++
		//fmt.Printf("isPreifx: %v, len: %d\n", is_prefix, len(data))
		fmt.Printf("%2d: %s\n", n, data)
	}
}

func main() {
	//test1()
	test2()
}
