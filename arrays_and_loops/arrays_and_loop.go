package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [5]int
	for i := range arr {
		arr[i] = i*i
		var odd string
		if arr[i]%2==1 {
			odd="odd"
		} else {
			odd="even"
		} 
		fmt.Printf("Element %d: %d , Square %#v , %s", i+1, arr[i], math.Sqrt(float64(arr[i])), odd+"\n") 
	}
}
