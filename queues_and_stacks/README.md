# Assignment 13: Stacks and Queues Combined

## Objective

This assignment introduces **stacks** and **queues** and demonstrates how to use them together by implementing a program that processes data in both **First-In-First-Out (FIFO)** and **Last-In-First-Out (LIFO)** order.

## Definitions

- **Queue**: A queue is a data structure that follows the First-In-First-Out (FIFO) principle, meaning elements are added at the end and removed from the front.
- **Stack**: A stack is a data structure that follows the Last-In-First-Out (LIFO) principle, where elements are added and removed from the top.

## Instructions

1. Create a file named `stack_queue_assignment.go`.
2. Inside the file, write a Go program that:
   - Defines a `Queue` struct with `Enqueue` (add to end) and `Dequeue` (remove from front) methods.
   - Defines a `Stack` struct with `Push` (add to top) and `Pop` (remove from top) methods.
   - Prompts the user to enter a series of integers.
   - Adds each integer to both a queue and a stack.
   - Once all integers are added, it:
     - Dequeues elements from the queue and prints them in FIFO order.
     - Pops elements from the stack and prints them in LIFO order.
3. Update the README with instructions on how to compile and run this program.

## Example Scenario

- **Input**: 1, 2, 3, 4
- **Queue Output (FIFO)**: 1, 2, 3, 4
- **Stack Output (LIFO)**: 4, 3, 2, 1

**Example Output:**

```text
Enter integers (separated by space): 1 2 3 4
Queue output (FIFO): 1, 2, 3, 4
Stack output (LIFO): 4, 3, 2, 1
```
