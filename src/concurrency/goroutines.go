package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func test1() {
	go say("world")
	go say("hello")
	say("haha...")
}

func main() {
	test1()
}
