package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error: Invalid number of arguments")
		os.Exit(1)
	}

	// 第一引数を整数に変換
	result, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to convert argument to integer")
		os.Exit(1)
	}

	// アセンブリコードの出力
	fmt.Println(".global main")
	fmt.Println("main:")
	fmt.Printf("	mov x0, #%d\n", result)
	fmt.Println("	ret")
}
