package main

import "fmt"

func main() {
        m := make(map[string]int, 3)
        m["apple"] = 1
        m["oranges"] = 4
        m["bananas"] = 11

        m["bananas"] += 1
        m["oranges"] -= 1

        for k, v :=range(m) {
                fmt.Printf("\n Item: %s , Quantity: %d", k, v)
        }
}
