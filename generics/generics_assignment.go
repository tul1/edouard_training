package main

import (
	"fmt"
	"strings"
)

type Person[T int | string] struct {
	Name string
	Age T
}

func FindMax[V int | float64](slice []V) V {
	var maxval V
	for _, v := range slice[1:] {
		if v>maxval {
			maxval=v
		}
	}
	return maxval
}

func FindOldest[T int | string](slice []Person[T]) { // takes a slice of Person struct with Age attribute constrain by type T
	if len(slice) == 0 {
		fmt.Println("No people in the slice.")
		return
	}

	oldest := slice[0] // Start with the first person.

	// Use a type switch to handle the two types of T.
	switch any(oldest.Age).(type) {
	case int:
		for _, person := range slice {
			if any(person.Age).(int) > any(oldest.Age).(int) { // cause type assertion need value of an interface, i attached type parameter to an interface (any)
				oldest = person 
			}
		}
		fmt.Printf("Oldest Person (numeric): %s, Age: %v", oldest.Name, oldest.Age)

	case string:
		for _, person := range slice {
			if strings.Compare(any(person.Age).(string), any(oldest.Age).(string)) > 0 {
				oldest = person
			}
		}
		fmt.Printf("Oldest Person (lexicographic): %s, Age: %v\n", oldest.Name, oldest.Age)

	/**default:
		fmt.Println("Unsupported type for Age.")
	} Age of a Person struct is forcially constrained by type T and type T include only next underlying values : int and string**/
	}
}

func main() {

	arrFindMaxInt := [3]int{4, 8, 13}
	fmt.Println("Max of int slice:",FindMax(arrFindMaxInt[:]))

	arrFindMaxFloat := [3]float64{4.2, 4.1, 4.3}
	fmt.Println("Max of float slice:",FindMax(arrFindMaxFloat[:]))

	arrFindOldestInt := []Person[int]{Person[int]{Name: "person1", Age: 102}, Person[int]{Name: "person2", Age: 210}}
	FindOldest(arrFindOldestInt)

	fmt.Println()

	arrFindOldestString := []Person[string]{Person[string]{Name: "person1", Age: "10"}, Person[string]{Name: "person2", Age: "12"}}
	FindOldest(arrFindOldestString)

}

