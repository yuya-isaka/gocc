package main

import (
	"fmt"
	"os"
	"strconv"
)

func atoi(rn rune) int {
	result, err := strconv.Atoi(string(rn))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to convert argument to integer")
		os.Exit(1)
	}
	return result
}

func isDigit(rn rune) bool {
	return '0' <= rn && rn <= '9'
}

func strlen(p []rune, idx *int) int {
	tmp := 0
	for *idx < len(p) && isDigit(p[*idx]) {
		tmp = (tmp * 10) + atoi(p[*idx])
		*idx++
	}
	return tmp
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error: Invalid number of arguments")
		os.Exit(1)
	}

	p := []rune(os.Args[1])
	idx := 0

	// アセンブリコードの出力
	fmt.Println(".global main")
	fmt.Println("main:")
	fmt.Printf("	mov x0, #%d\n", strlen(p, &idx))

	for idx < len(p) {
		switch p[idx] {
		case '+':
			idx++
			fmt.Printf("	add x0, x0, #%d\n", strlen(p, &idx))
		case '-':
			idx++
			fmt.Printf("	sub x0, x0, #%d\n", strlen(p, &idx))
		default:
			break
		}
	}

	fmt.Println("	ret")
}
