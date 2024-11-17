
# Assignment 12: Linked Lists

## Objective

This assignment introduces the concept of a **linked list** and demonstrates how to create and manage nodes within it.

## Definitions

- **Linked List**: A linked list is a data structure consisting of nodes where each node contains data and a reference (or link) to the next node in the sequence. It is useful for dynamic memory allocation and when the list size may change frequently. Unlike arrays, linked lists do not require contiguous memory.

## Instructions

1. Create a file named `linkedlist_assignment.go`.
2. Inside the file, write a Go program that:
   - Defines a `Node` struct with two fields: `value` (int) and `next` (a pointer to the next `Node`).
   - Defines a `LinkedList` struct with a `head` pointer to the first `Node`.
   - Implements methods to add a new node to the end of the list and to print all node values sequentially.
   - Demonstrates creating a linked list, adding nodes, and printing the linked list’s values.
3. Update the README with instructions on how to compile and run this program.

**Example Output:**

```text
Linked List: 3 -> 5 -> 7
```

## How it works

1- Declare Node struct with next attributes: value of type int and next of type pointer (directed to a Node struct)

2- Declare LinkedList struct with head attribute, head is a pointer type (directed to Node struct)

3- Create a Add LinkedList medthod, it takes data of type int

3.1- create a unperson pointer, it takes data as value and nil as next

3.2- if LinkedList is empty then create a node with unperson pointer else create a temporary pointer and assign Linkedlist to it 

3.3- Pass through each node of temporary pointer with next attribute, check at each iteration that next attribute value of 
temporary pointer is not nil

3.4- if next attribute of temporary pointer is not nil then assign actual pointer to next pointer of temporary variable, if next 
attribute of actual temporary pointer is nil then add an unperson node as next attribute of actual temporary pointer

## How it runs 

Execute go code:

```bash
go run linkedlist_assignment.go
```

Compilation:

```bash
go build linkedlist_assignment.go
```

Execute binary:

```bash
./linkedlist_assignment
```

## Output

output:
LinkedList: <LinkedList.value> -> <LinkedList.next.value>
