# Assignment 5: Pointers, Structs, and Pointer Semantics

## Objective

This assignment will introduce pointers, structs, and the difference between **pointer semantics** and **value semantics** in Go. You will learn how to pass values by pointer to modify data directly and how to pass values by value, which creates a copy of the data.

## Definitions

- **Pointers**: A pointer is a variable that holds the memory address of another variable. It allows functions to modify the actual variable by pointing to its memory location instead of working with a copy of the value.

- **Pointer Semantics**: When a variable is passed using pointer semantics, the function operates on the actual memory location of the variable, allowing changes to persist outside of the function. This is efficient when passing large data structures like arrays or structs.

- **Value Semantics**: When a variable is passed by value, a copy of the variable is made. Changes to the copy do not affect the original variable. This is the default behavior in Go when passing data types like integers or structs unless pointers are explicitly used.

## Instructions

1. Create a file named `pointers_structs.go`.
2. Inside the file, write a Go program that:
   - Defines a `Rectangle` struct with fields `width` and `height` (both of type `float64`).
   - Creates a function `ScaleRectByPointer` that takes a pointer to a `Rectangle` and a scaling factor. It modifies the rectangle’s dimensions using pointer semantics.
   - Creates another function `ScaleRectByValue` that takes a `Rectangle` by value (no pointer) and a scaling factor. It modifies the copy of the rectangle’s dimensions using value semantics.
   - Demonstrates both pointer and value semantics by creating a `Rectangle` instance, scaling it with both functions, and printing the rectangle’s dimensions before and after each function call.
   - Highlights the difference in behavior: the changes from `ScaleRectByPointer` persist, while those from `ScaleRectByValue` do not.
3. Update the README with instructions on how to compile and run this program.

**Example Output:**

```
Original Rectangle: Width = 5, Height = 3
Scaled by Value (No Change): Width = 5, Height = 3
Scaled by Pointer (Changed): Width = 10, Height = 6
```
