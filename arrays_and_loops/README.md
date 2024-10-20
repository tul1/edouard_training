# Assignment 2: Arrays and Loops

## Objective

The goal of this assignment is to practice working with arrays and loops in Go. It introduces basic data structures, iteration, and conditional logic to manipulate and display values within an array.

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

Array Squaring and Even/Odd Identifier

This program demonstrates the use of arrays, loops, and basic mathematical operations in Go. It initializes an array of integers, computes the square of each element, and 
checks whether the number is odd or even. Additionally, the program prints the square root of each element for reference.
How It Works

    An array of integers with 5 elements is created.
    A for loop iterates over each element of the array:
        The square of the index is assigned to the respective array element.
        The program checks whether the number is odd or even.
        It prints the following information for each array element:
            The element’s position in the array.
            The squared value.
            The square root of the element.
            Whether the number is odd or even.

Output Format

For each element in the array, the program prints the following formatted output:

Element <position>: <value>, Square: <square root>, <odd/even>

