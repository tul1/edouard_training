# Assignment 4: Structs, Public and Private Methods

## Objective

This assignment introduces structs and how to create public and private methods in Go.

## Instructions

1. Create a file named `structs_assignment.go`.
2. Inside the file, write a Go program that:
   - Defines a `Person` struct with fields `firstName` (string), `lastName` (string), and `age` (int).
   - Creates a private method `isAdult()` that checks if a person is 18 years or older.
   - Creates a public method `PrintInfo()` that prints the person’s full name and whether they are an adult (using the `isAdult()` method).
3. Update the README with instructions on how to run this program.

**Example Output:**

```
Full Name: John Doe, Age: 20
John is an adult.
```

## How it works



1- Initialize Person struct with firstName, lastName string & age int



2- Private Person method check if age is less than 18 & return a bool



3- Public PrintInfo check if p1 method is true or false



## How it runs



Compilation: `go build structs_public_private_methods.go`



Execute go code: `go run structs_public_and_private_methods.go`



Execute binary: `./structs_public_and_private_methods`



## Output



Formated output:

<Person.firstName> <Person.isAdult()>

Full Name: <Person.firstName> <Person.lastName> , age: <Person.age>
