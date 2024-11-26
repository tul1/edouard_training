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

## How it works

1- Declare TreeNode struct and its attributes : value of type int, left of type pointer (directed to a TreeNode struct), right 
of type pointer (directed to a TreeNode struct)

2- Declare BinaryTree struct and its attributes : root of type pointer (directed to a TreeNode struct)

3- Add a BinaryTree method named Insert and takes Data argument of type int

3.1- declare temp_pointer and define it to be attached to the root pointer of a BinaryTree struct

3.2- Declare a if statement to check if BinaryTree root pointer exist, if not then create a TreeNode struct and assign it to 
BinaryTree.root pointer

3.3- if BinaryTree root pointer exist then check if data value (argument of method) is equal, less or more than 
BinaryTree.root.value

3.4- if data value is higher than BinaryTree.root.value then iterate over each TreeNode of the BinaryTree with right pointer 
attribute of each TreeNode and check at each iteration if TreeNode value is equal to data, if true then assign data value to 
TreeNode else continue iteration
At the end of iteration, if value is not found then create a new TreeNode to the last BinaryTree.right

3.5- if data value is equal to BinaryTree.root.value then assign data value to BinaryTree.root.value var


3.6- if data value is less than BinaryTree.root.value then iterate over each TreeNode of the BinaryTree with left pointer 
attribute of each TreeNode and check at each iteration if TreeNode value is equal to data, if true then assign data value to 
TreeNode else continue iteration
At the end of iteration, if value is not found then create a new TreeNode to the last BinaryTree.left

4- Declare a BinaryTree struct

4.1- Use Insert Method of BinaryTree struct to add new value

4.2- declare and define a temp pointer related to BinaryTree root pointer in order to print each TreeNode value in ascending 
order 

## How it runs 

Execute go code:

```bash
go run binary_tree_assignment.go
```

Compilation:

```bash
go build binary_tree_assignment.go
```

Execute binary:

```bash
./binary_tree_assignment
```

## Output

output:
In-order traversal: <BinaryTree.(root, left or right).value>

