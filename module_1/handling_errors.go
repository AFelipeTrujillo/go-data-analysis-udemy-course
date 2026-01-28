package main

import (
	"fmt"
	"strconv"
)

func convertStringToInt(s string) (int, error) {
	// This function is intentionally left blank.
	number, err := strconv.Atoi(s)
	return number, err
}

func main() {
	// This file is intentionally left blank.
	defer fmt.Println("Execution completed.")
	result, err := convertStringToInt("123")
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Converted number:", result)
	}

	SayHello := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}

	SayHello("Alice")

}
