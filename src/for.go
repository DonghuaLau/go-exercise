package main

import "fmt"

func test1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func test2() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

// for as while
func test3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// loop forever
func test4() {
	n := 0
	for {
		n++
		fmt.Println(n)
	}
}

func main() {
	test1()
	test2()
	test3()
	//test4()
}

