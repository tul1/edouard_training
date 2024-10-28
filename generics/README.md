# Assignment 11: Generics in Go with Generic Field Types

## Objective

This assignment expands on **generics** in Go by allowing a struct field to use a generic type, showcasing how generics can make data structures more flexible.

## Definitions

- **Generics**: Generics allow defining types and functions that work with any data type, using type parameters. This assignment demonstrates defining struct fields with a generic type, enabling flexible data handling and operations.

## Instructions

1. Create a file named `generics_assignment.go`.
2. Inside the file, write a Go program that:
   - Defines a generic function `FindMax` to return the maximum value in a slice of any ordered type (e.g., integers, floats).
   - Defines a struct `Person` with fields `Name` (string) and `Age` as a generic type, `T`.
   - Creates a function `FindOldest` to find the oldest person in a slice of `Person` structs, handling `Age` generically for either `int` or `string` types.
   - Demonstrates the `FindMax` function with slices of integers and floats, printing the maximum values.
   - Demonstrates the `FindOldest` function with two slices of `Person` structs: one with `Age` as `int` and one with `Age` as `string`.
3. Update the README with instructions on how to compile and run this program.

**Example Output:**

```text
Max of int slice: 9
Max of float slice: 5.67
Oldest person (int age): Carol, Age: 40
Oldest person (string age): Alice, Age: 45
```
