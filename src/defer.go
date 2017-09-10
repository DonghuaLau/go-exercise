package main

import "fmt"

func test1() {
	n := 1
	defer fmt.Printf("world, %d\n", n)

	fmt.Printf("hello, %d\n", n)
}

func test2() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	test1()
	test2()
}
