package main

import "fmt"

type Person struct {
        firstName, lastName string
        age uint
}

func (p Person) isAdult() bool {
        return p.age < 18
}

func PrintInfo(p1 Person) {
        if p1.isAdult()==true {
                fmt.Printf("%s isn't an adult\n", p1.firstName)
        } else {
                fmt.Printf("%s is an adult\n", p1.firstName)
        }    
        fmt.Printf("Full Name: %s %s, age: %d\n", p1.firstName, p1.lastName, p1.age)        
}

func main() {
        p1 := Person{"John", "Doe", 21}
        PrintInfo(p1)   
}
             
