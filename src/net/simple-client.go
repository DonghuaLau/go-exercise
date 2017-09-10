package main

import (
	"fmt"
	"net"
	"bufio"
)

func test1() {
	conn, err := net.Dial("tcp", "liudonghua.net:80")
	if err != nil {
		// handle error
		fmt.Println("connect failed")
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: liudonghua.net\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	//buf := make([]byte, 1024) 
	
	fmt.Println(status)

}

func main() {
	test1()
}
