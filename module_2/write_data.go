package module2

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func WriteData() {

	currentTime := time.Now()
	year := currentTime.Year()

	file, err := os.Open("./module_2/dataset2.csv")
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

	outFile, err := os.Create("./module_2/dataset3.csv")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	// what is flush?
	// Ensures that all buffered data is written to the underlying writer.
	defer writer.Flush()

	newHeader := append(header, "birth_year")
	if err := writer.Write(newHeader); err != nil {
		panic(err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		ageStr := record[3]
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Println("Error converting age:", err)
			continue
		}

		birthYear := fmt.Sprintf("%d", (year - age))
		newRecord := append(record, birthYear)

		if err := writer.Write(newRecord); err != nil {
			panic(err)
		}
	}

}
