package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [5]int
	for i:=0;i<len(arr);i++ {
		arr[i] = i*i
		var odd string
		if arr[i]%2==1 {
			odd="odd"
		} else {
			odd="even"
		} 
		fmt.Print("\nElement ", i+1, ": ", arr[i], ",  Square: ", math.Sqrt(float64(arr[i])), ", ",  odd, "\n") 
	}
}
