package main

import (
	"com/stringutil"
	"fmt"
)

func test1() {
	fmt.Printf("Another main enter\n")
	fmt.Println(stringutil.Reverse("!oG, olleH"))
}

func main() {
	test1()
	test2()
}
