package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("os.Argsが2ちゃうで")
	}
	fmt.Println(".global main")
	fmt.Println("main:")
	result, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Stderr.WriteString("Atoi失敗")
	}
	fmt.Printf("	mov x0, #%d\n", result)
	fmt.Println("	ret")
}
