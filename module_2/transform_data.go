package module2

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func TransformData() {
	// Placeholder function for data transformation logic

	file, err := os.Open("./module_2/dataset2.csv")
	currentTime := time.Now()
	year := currentTime.Year()

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()

	if err != nil {
		panic(err)
	}
	fmt.Println("Header:", header)

	fmt.Println("Transforming data...")

	for {
		record, err := reader.Read()

		if err != nil {
			break
		}

		// fmt.Println("Original Record:", record)

		id := record[0]
		firstNameStr := record[1]
		// lastNameStr := record[1]
		ageStr := record[3]
		// emailStr := record[3]
		// genderStr := record[4]
		// ipAddressStr := record[5]

		age, err := strconv.Atoi(ageStr)

		if err != nil {
			fmt.Println("Error converting age:", err)
			continue
		}

		if age < 30 {
			continue
		}

		birthYear := year - age

		fmt.Printf("ID: %s, First Name: %s, Birth Year: %d\n", id, firstNameStr, birthYear)

	}

}
