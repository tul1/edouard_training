package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
	var stdin string
	fmt.Scan(&stdin)
	fmt.Println("Original String:",stdin)
	fmt.Println("Uppercase  String:", strings.ToUpper(stdin)) 
	fmt.Println("Reversed String:", Reverse(stdin))
}
