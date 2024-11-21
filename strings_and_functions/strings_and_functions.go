package main

import (
        "fmt"
        "strings"
)

func Reverse(s string) string {
    byteR := []byte(s)
    for i, j := 0, len(byteR)-1; i < j; i, j = i+1, j-1 {
        byteR[i], byteR[j] = byteR[j], byteR[i]
    }
    return string(byteR)
}

func main() {
        var stdin string
        _,err:=fmt.Scan(&stdin)
        if err!= nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Original String: %s , Uppercase String: %s , Reversed String: %s", stdin, strings.ToUpper(stdin),Reverse(stdin))
	}
}


