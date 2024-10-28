package main

import "fmt"

type Rectangle struct {
        width, height float64
}

func (r1 *Rectangle) ScaleRectByPointer() {
        r1.height = 1.6
        r1.width = 1.6
}

func (r1 Rectangle) ScaleRectByValue() {
        r1.height = 1.6
        r1.width = 1.6
}

func main() {
        r1 := Rectangle{3, 3}
        fmt.Printf("Original Rectangle: Width = %.1f , Height = %.1f", r1.width, r1.height)

        r1.ScaleRectByValue()
        fmt.Printf("Scaled by Value (No Change): Width = %.1f , Height = %.1f", r1.width, r1.height)

        r1.ScaleRectByPointer()
        fmt.Printf("Scaled by Pointer (Changed): Width = %.1f , Height = %.1f \n", r1.width, r1.height)
}
