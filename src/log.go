package main

import (
	"fmt"
	"os"
	"bytes"
	"log"
)

func test1() {
	//log.Print("call log.Print")
	//log.Fatal("call log.Fatal")
	log.Panic("call log.Panic")
	//panic("call panic")
}

func test2() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "[logger] ", log.Lshortfile | log.Ldate | log.Ltime)
	)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
}

func test3() {
	var (
		logfile, err = os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		buf    bytes.Buffer
		logger = log.New(&buf, "[logger] ", log.Lshortfile | log.Ldate | log.Ltime)
	)

	if err != nil {
		log.Fatal(err)
	}

	logger.SetOutput(logfile)

	logger.Print("Hello, log file!")
	logger.Fatal("Hey, fatal log")

	//fmt.Print(&buf)
}


func main() {
	//test1()
	//test2()
	test3()
}
