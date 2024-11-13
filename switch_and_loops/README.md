# Assignment 7: Switch and Loops

## Objective

This assignment will help you understand control flow using logic operators in Go, focusing on `switch`, `while`, and `do while` equivalents (`for` loop in Go).

## Instructions

1. Create a file named `logic_operators.go`.
2. Inside the file, write a Go program that:
   - Reads an integer from user input.
   - Uses a `switch` statement to print a message depending on the value of the number (e.g., "positive", "negative", or "zero").
   - Uses a `for` loop (simulating a `while`) to count down from the entered number to zero and print each value.
   - Uses a `for` loop (simulating a `do-while`) that keeps asking the user to input a number until they enter a positive integer.
3. Update the README with instructions on how to compile and run this program.

**Example Output:**

```
Enter a number: 3
The number is positive.
Counting down: 3, 2, 1, 0
Enter a number: -1
Please enter a positive number: -5
Please enter a positive number: 2
You entered a positive number: 2
```

## How it works

1- Scan user input and assign value into a space of memory

2- In case of user input is positive print it and start a loop from user value to 0, at each iteration print value of user input

3- In case of user input is negative ask user to change he's input to a positive value with 

4- Use default case in midle switch element in order to help iteration memory

## How it runs 

Execute go code:

```bash
go run logic_operators.go
```

Compilation:

```bash
go build logic_operators.go
```

Execute binary:

```bash
./logic_operators
```

## Output

output:
The number is positive
formated output:
Counting down: <user_input>
output:
Please enter a positive number:
formated output:
You entered a positive number: <user_input>
output:
The number is zero
