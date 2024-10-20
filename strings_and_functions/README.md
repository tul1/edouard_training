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

# Exercie 3

##

Create a user input and manpulate it with runes(int32) to rend initial string, string upcase, reversed string

## How it works

1- Initialize var stdin type string
2- create user input and attach value to var stdin
3- display initial string with Println func of fmt package
4- display string upcase with ToUpper func from strings package
5- display reversed string call Reverse(return string) func with stdin parameter 
5a -create slic of int32
5b - create loop and initialize i and j var that are -1 and 1 
5c - iterate over loop in order to replace each 'index/position' with opposite
