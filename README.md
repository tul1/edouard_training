# Assignment 2: Arrays and Loops

## Objective

The goal of this assignment is to practice working with arrays and loops in Go. It introduces basic 
data structures, iteration, and conditional logic to manipulate and display values within an array.

## Instructions

1. Create a file named `arrays_assignment.go`.
2. Inside the file, write a Go program that:
   - Declares an array of integers with 5 elements.
   - Uses a loop to populate the array with values.
   - Uses a second loop to print each element, its square, and whether the element is odd or even.
3. Update the README with instructions on how to compile and run this program.

**Example Output:**

```
Element 1: 1, Square: 1, Odd
Element 2: 2, Square: 4, Even
Element 3: 3, Square: 9, Odd
Element 4: 4, Square: 16, Even
Element 5: 5, Square: 25, Odd
```

compil : go build arrays_assignment.go
run : go run arrays_assignment.go
exec : ./arrays_assignment.go


package main // define package

import ( // import packages
	"fmt"
	"math"
)

func main() { // func of beginning
var arr [5]int // allocate an array of 5int to the arr var
for i:=0;i<len(arr);i++ { // loop that iterate on var lenght
arr[i] = i*i // stock like Pow Func exec of a value in the array relaed to var arr
var ode string  

if(arr[i]%2==1){ode="odd"}else{ode="even"} // statement to verify if modulo of the value  is odd or 
even

fmt.Printf(
"Element %v : %v , Square : %v , %v", i+1, // % is used to desribe the struct/type format that we want 
to recover with text, depends of  value
arr[i], math.Sqrt(float64(arr[i])) ,ode + " \n") // return value of index i in arr array
}
}


