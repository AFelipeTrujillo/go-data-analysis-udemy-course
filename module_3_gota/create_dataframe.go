package module3

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gota/gota/dataframe"
)

func CreateDataframe() {

	file, err := os.Open("./datasets/dataset2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)

	log.Println("Print Data Frames Titles")
	log.Println(df.Names())
	fmt.Println(df)

	num_rows := df.Nrow()
	num_cols := df.Ncol()

	log.Println("Row Nums", num_rows)
	log.Println("Row Cols", num_cols)

	csvString := `name,age,year
andy,37,1988
joha,37,1988`

	df_from_string := dataframe.ReadCSV(strings.NewReader(csvString))
	fmt.Println(df_from_string)

}
