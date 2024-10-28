# Assignment 8: Introduction to Context and Goroutines

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

## How it works

1- define a context withdeadline with a duration of 2seconds, when it starts it takes actual time and add 2seconds

2- execute Cancelfunc of context if main func is terminated to not let context survive more that what is expected

3- execute a go routine to print a message, wait 5seconds et print another message

4- From context.Done channel if deadline expires so print Error method of
context 

## How it runs

Execute go code:
   
```bash
go run context_goroutines_assignment.go
```

Compilation:
   
```bash
go build context_goroutines_assignment.go
```

Execute binary:

```bash
./context_goroutines_assignment
```

## Output

output:
context deadline exceeded
