package main

import "fmt"

func main() {
	// This file is intentionally left blank.

	var age int = 30

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Minor")
	}

	if age := 25; age >= 18 {
		fmt.Println("Adult with scoped age:", age)
	}

	// Loops
	// FOR LOOP
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// WHILE LOOP (using for)
	j := 0
	for j < 5 {
		fmt.Println("While Iteration:", j)
		j++
	}

	// RANGE LOOP
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	users := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
		"Diana":   28,
	}
	for name, age := range users {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

}
