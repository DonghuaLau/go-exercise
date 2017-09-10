package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
)

func main() {
	
	name := "Tom"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	conn, err := net.Dial("tcp", ":18080")
	if err != nil {
		// handle error
		fmt.Println("connect failed")
		return
	}
	fmt.Fprintf(conn, name)
	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(status)
}
