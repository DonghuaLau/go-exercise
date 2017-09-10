package main

import (
	"os"
	"fmt"
)

func main() {

    args := os.Args
    args_len := len(args)

    arg := args[3]

    fmt.Println(args_len, args)
    fmt.Println(arg)
}

