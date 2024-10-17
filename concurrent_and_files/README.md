# Assignment 6: Concurrent File Writing with Goroutines

## Objective

This assignment reinforces goroutines and synchronization when multiple goroutines write to a file concurrently.

## Instructions

1. Create a file named `file_writer_assignment.go`.
2. Inside the file, write a Go program that:
   - Starts 3 goroutines, each writing a unique message to the same file.
   - Uses synchronization mechanisms (like channels or `sync.WaitGroup`) to ensure all goroutines finish writing before the program exits.
3. Update the README with instructions on how to run this program.

**Example Output (from the file):**

```
Message from Goroutine 1
Message from Goroutine 2
Message from Goroutine 3
```
