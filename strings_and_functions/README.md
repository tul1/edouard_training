# Assignment 3: String Manipulation

## Objective

This assignment introduces working with strings and string functions in Go, focusing on manipulating and transforming strings.

## Instructions

1. Create a file named `strings_assignment.go`.
2. Inside the file, write a Go program that:
   - Reads a string from the user input (e.g., "Hello, Gophers!").
   - Converts the string to uppercase and prints the result.
   - Reverses the string and prints the reversed version.
3. Update the README with instructions on how to run this program.

**Example Output:**

```
Original String: Hello, Gophers!
Uppercase String: HELLO, GOPHERS!
Reversed String: !srehpoG ,olleH
```

## How it works

1- Capture user input as string

2- Extract char from input and insert them into a slice of byte type
   
3- Loop over slice with i & j until they have same index, i & j are first and last value of slice
   
## How it runs

Compilation: `go build strings_and_functions.go`

Execute go code: `go run strings_and_functions.go`

Execute binary: `./strings_and_functions`

## Output

Formated output:
Original String: <stdin(I/O)> , Uppercase String: <strings.ToUpper(stdin)> , Reversed String: <Reverse(stdin)>
