# Assignment 17: Advanced Pointers, Arrays, and Matrices

## Objective

This assignment provides an in-depth look at pointers, arrays, and matrices in Go, focusing on both 1D and 2D arrays and utilizing pointers for in-place data manipulation. By the end, you will have implemented a variety of matrix operations, including element-wise updates, summing rows and columns, matrix addition, and calculating the determinant.

## Definitions

- **Pointer**: A pointer holds the memory address of another variable, allowing direct data manipulation at that location.
- **Array**: An array is a fixed-size collection of elements of a single type. In Go, arrays are indexed and their size is defined at declaration.
- **Matrix**: A matrix is a 2D array that organizes data in rows and columns, with each element accessed via two indices.

## Instructions

1. **Create a file** named `advanced_pointers_arrays_matrices.go`.
2. Inside the file, implement the following tasks to perform complex operations on arrays and matrices using pointers.

## Tasks

1. **Array Operations**:
   - Declare an array of 10 integers.
   - Write a function `initializeArray` that takes a pointer to the array and populates it with the squares of its indices (i.e., `0, 1, 4, 9, ...`).
   - Write a function `reverseArray` that reverses the array in place using pointers.
   - In `main`, print the array after initialization and after reversal.

2. **Matrix Operations**:
   - Declare a 3x3 integer matrix and initialize it with values from 1 to 9.
   - Write a function `sumRows` that takes a pointer to the matrix and returns a slice with the sum of each row.
   - Write a function `sumColumns` that takes a pointer to the matrix and returns a slice with the sum of each column.
   - In `main`, print the original matrix, row sums, and column sums.

3. **Matrix Addition**:
   - Declare another 3x3 matrix initialized with values from 9 to 1.
   - Write a function `addMatrices` that takes pointers to two 3x3 matrices and returns a new 3x3 matrix that is the element-wise sum of the two matrices.
   - In `main`, print the result of the matrix addition.

4. **Transpose and Determinant Calculation**:
   - Write a function `transposeMatrix` to transpose a 3x3 matrix in-place, swapping rows and columns.
   - Write a function `calculateDeterminant` that takes a pointer to a 3x3 matrix and returns its determinant. (For a 3x3 matrix, use the formula:
     \[
     \text{det} = a(ei - fh) - b(di - fg) + c(dh - eg)
     \]
     assuming the matrix is represented as:
     ```
     [a, b, c]
     [d, e, f]
     [g, h, i]
     ```
   - In `main`, print the transposed matrix and its determinant.

## Example Output

```
Initialized Array: [0 1 4 9 16 25 36 49 64 81]
Reversed Array: [81 64 49 36 25 16 9 4 1 0]

Original Matrix:
1 2 3
4 5 6
7 8 9

Row Sums: [6, 15, 24]
Column Sums: [12, 15, 18]

Matrix Addition Result:
10 10 10
10 10 10
10 10 10

Transposed Matrix:
1 4 7
2 5 8
3 6 9

Determinant: 0
```
