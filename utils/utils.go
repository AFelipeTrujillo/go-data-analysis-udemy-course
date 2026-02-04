package utils

import (
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func CreateDataFrame(csvFilePath string) dataframe.DataFrame {
	file, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return dataframe.ReadCSV(file)
}
