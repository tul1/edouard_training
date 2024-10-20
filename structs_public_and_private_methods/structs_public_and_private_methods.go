package main

import "fmt"

type Person struct {
        firstName, lastName string
        age uint
}

func (p Person) isAdult() (uint, string, string) {
        return p.age, p.firstName, p.lastName
}

func PrintInfo(p1 Person) {
        p1a, p1f, p1l := p1.isAdult()
        switch {
                case p1a < 18:
                        fmt.Printf("\n Full Name: %v  %v, Age: %v \n %v %v isn't an adult\n\n", p1f, p1l, p1a, p1f, p1l)
                case p1a >= 18:
                        fmt.Printf("\n Full Name: %v  %v, Age: %v \n %v %v is an adult\n\n", p1f, p1l, p1a, p1f, p1l)
                }

}

func main() {
p1 := Person{"John", "Doe", 18}
PrintInfo(p1)
}
