package main

import "fmt"

func main() {
        var user_input int
        fmt.Print("Enter a number: ")
        fmt.Scan(&user_input)
        switch {
        case user_input>0:
                fmt.Println("The number is positive")
                fmt.Printf("Counting down: %d ,", user_input)
                for user_input!=0 {
                        user_input-=1
                        fmt.Print(user_input, ", ")
                }
        default:
		fmt.Print("The number is zero")
	case user_input<0:		
                for user_input<0 {
                        fmt.Print("Please enter a positive number: ")
                        fmt.Scan(&user_input)
                }
                fmt.Printf("You entered a positive number: %d", user_input)
        }
}
