# Assignment 5: Introduction to Context and Goroutines

## Objective

This assignment introduces Go’s `context` package and demonstrates how to control goroutines with it.

## Instructions

1. Create a file named `context_goroutines_assignment.go`.
2. Inside the file, write a Go program that:
   - Starts a goroutine that simulates a task that takes time (e.g., sleeping for 5 seconds).
   - Uses the `context` package to allow for canceling the goroutine after 2 seconds.
   - Prints a message if the task is canceled due to the context.
3. Update the README with instructions on how to run this program.

**Example Output:**

```
Task started...
Task canceled due to context timeout!
```