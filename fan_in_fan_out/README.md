# Assignment 10: Concurrent File Writing with Fan-In and Fan-Out Goroutines

## Objective

This assignment introduces the **fan-out** and **fan-in** concurrency patterns by having multiple goroutines write messages to a file concurrently and consolidating the results.

## Definitions

- **Fan-Out**: Fan-out occurs when multiple goroutines are spawned to perform independent tasks concurrently. In this case, each goroutine will write to the file.
  
- **Fan-In**: Fan-in refers to merging the results of multiple goroutines into a single channel or result stream. Here, we will merge the status messages from each writing goroutine into a single channel.

## Instructions

1. Create a file named `file_writer_fanin_fanout.go`.
2. Inside the file, write a Go program that:
   - Spawns **three goroutines** (fan-out), where each goroutine writes a unique message to the same file.
   - Uses a **fan-in pattern** to collect status messages from the goroutines, indicating that the writing operation was successful for each one.
   - Writes the messages to the file concurrently but prints the success status messages in the main function using fan-in.
3. Update the README with instructions on how to compile and run this program.

## Additional Notes

- Ensure that access to the file is synchronized properly to avoid race conditions. You can use `sync.Mutex` or another synchronization mechanism.
- The main function should only print the success messages after receiving them from the fan-in channel.
  
## Example Output:

```
Task 1: Writing to file complete.
Task 3: Writing to file complete.
Task 2: Writing to file complete.
```
