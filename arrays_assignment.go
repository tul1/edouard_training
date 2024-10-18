package main

import (
	"fmt"
	"math"
)

func main() {
var arr [5]int
for i:=0;i<len(arr);i++ {
arr[i] = i*i
var ode string
if(arr[i]%2==1){ode="odd"}else{ode="even"}
fmt.Printf("Element %v : %v , Square : %v , %v", i+1, arr[i], 
math.Sqrt(float64(arr[i])) ,ode + " \n")
}
}

