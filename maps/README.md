# Assignment 6: Maps

## Objective

This assignment will help you learn Go’s data structures, particularly hash maps (called `map` in Go).

## Instructions

1. Create a file named `hashmap.go`.
2. Inside the file, write a Go program that:
   - Declares a map where the keys are strings (representing item names) and the values are integers (representing quantities).
   - Adds several items to the map (e.g., "Apples", "Oranges", "Bananas").
   - Updates the quantity of one item and removes another.
   - Loops through the map to print each item’s name and quantity.
3. Update the README with instructions on how to compile and run this program.

**Example Output:**

```
Item: Apples, Quantity: 10
Item: Oranges, Quantity: 5
```

## How it works

1- Initialize a m map with key of type string & value of type int

2- Define several element in map with key & value

3- Add & Remove values of a map with associate keys

4- Iterate over map with key & value

## How it runs

Execute go code:
```bash
go run hashmap.go
```

Compilation:
```bash
go build hashmap.go
```

Execute binary:
```bash
./hashmap
```

## Output

formated output:
Item: <m[key]> , Quantity: <m[value]>
