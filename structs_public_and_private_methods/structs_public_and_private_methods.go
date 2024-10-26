package main

import "fmt"

type Person struct {
        firstName, lastName string
        age int
}

func (p Person) isAdult() bool {
        return p.age < 18
}

func (p Person) PrintInfo() {
        if p.isAdult() {
                fmt.Printf("%s isn't an adult\n", p.firstName)
        } else {
                fmt.Printf("%s is an adult", p.firstName)
        }
    fmt.Printf("Full Name: %s %s, age: %d", p.firstName, p.lastName, p.age)
}


func main() {
        p := Person{"John", "Doe", 12}
        p.PrintInfo()
}
