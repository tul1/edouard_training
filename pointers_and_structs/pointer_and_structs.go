package main

import "fmt"

type Rectangle struct {
        width, height float64
}

func ScaleRectByPointer(r1 *Rectangle) {
        r1.height = 1.6
        r1.width = 1.6
}

func ScaleRectByValue(r1 Rectangle) {
        r1.height = 1.6
        r1.width = 1.6
}

func main() {
        r1 := Rectangle{3, 3}
        fmt.Printf("Original Rectangle: Width = %.1f , Height = %.1f", r1.width, r1.height)

        ScaleRectByValue(r1)
        fmt.Printf(" Scaled by Value (No Change): Width = %.1f , Height = %.1f", r1.width, r1.height)

        ScaleRectByPointer(&r1)
        fmt.Printf(" Scaled by Pointer (Changed): Width = %.1f , Height = %.1f \n", r1.width, r1.height)
}
