# Assignment 15: Binary Trees

## Objective

This assignment introduces the **binary tree** data structure, which is fundamental in sorting, searching, and hierarchical data storage.

## Definitions

- **Binary Tree**: A binary tree is a data structure in which each node has at most two children, often referred to as the left and right child. Binary trees are commonly used for efficient searching and sorting. Special types, like binary search trees (BSTs), follow specific ordering rules to maintain order.

## Instructions

1. Create a file named `binary_tree_assignment.go`.
2. Inside the file, write a Go program that:
   - Defines a `TreeNode` struct with `value` (int) and pointers to `left` and `right` child nodes.
   - Defines a `BinaryTree` struct with a `root` pointer.
   - Implements methods to insert values following binary search tree rules and to traverse the tree in-order (left, root, right).
   - Demonstrates adding elements to the tree and printing them in ascending order.
3. Update the README with instructions on how to compile and run this program.

**Example Output:**

```text
Inserting values: 5, 3, 7, 2, 4
In-order traversal: 2 3 4 5 7
```
