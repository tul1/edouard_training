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

# Exercie 4

Create a character named John Doe aged of 18 and display if is an adult or not

## How it works

1- Create Person struct that take Firstname, Lastname var attribute that are string and one other parameter that is age and type of uint
2- Create isAdult private method that is attach to Person Object and return value of type bool depend if var.age is > 18 or not
3- Create public PrintInfo func that implement p1 var of type Person
3a- declare statement to know if p1 age atribute is > 18 or not, if it is then print it else print too
3b- Display format string with p1 next values (firstname, lastname, age)
4- Initialize p1 var that is type of Person and allocate it next values : "John", "Doe", 18
4- start PrintInfo func with p1 var
