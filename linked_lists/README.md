
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
