package module2

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ValidateData() {

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

	fmt.Println(header)

	headerLen := len(header)

	for {

		record, err := reader.Read()

		if err != nil {
			break
		}

		if len(record) < headerLen {
			fmt.Println("The length is not correct")
			continue
		}

		ageStr := record[3]

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			log.Println("Integer convertation error")
			continue
		}

		age++

		// Ingnoring blank lines
		if len(record) == 0 {
			continue
		}

		log.Println("OK Record")

	}

}
