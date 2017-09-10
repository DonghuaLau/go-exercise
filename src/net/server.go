package main

import (
	"fmt"
	"net"
	"bufio"
)

func handleConnection(conn net.Conn) {
	fmt.Println("new connection")

	buf := make([]byte, 1024)
	reader := bufio.NewReader(conn)
	_, err := reader.Read(buf)
	//status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("read failed, error: ", err)
		return
	}

	fmt.Printf("Read: %s\n", buf)
	
	fmt.Fprintf(conn, "Hi %s", buf)

	conn.Close()
}

func main() {

	ln, err := net.Listen("tcp", ":18080")
	if err != nil {
		// handle error
		fmt.Println("listen failed, error: ", err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("accept failed, error: ", err)
			return
		}
		go handleConnection(conn)
	}
}

