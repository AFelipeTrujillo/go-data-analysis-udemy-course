package module2

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV() {

	file, err := os.Open("./module_2/dataset.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		fmt.Println(record)
	}
}
