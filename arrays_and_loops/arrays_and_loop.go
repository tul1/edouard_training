package main

import (
        "fmt";
        "math";
)


func main() {
        var arr [5]int
        var pairity string
        for i := range arr {
                arr[i] = i*i
                if arr[i]%2==1 {
                        pairity="odd"
                } else {
                        pairity="even"
                }
                fmt.Printf("Element %d: %d , Square %#v , %s %s", i+1,
arr[i], math.Sqrt(float64(arr[i])), pairity, "\n")
        }
}

