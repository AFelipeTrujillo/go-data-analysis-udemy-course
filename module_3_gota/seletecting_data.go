package module3

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func SelectingData() {

	file, err := os.Open("./datasets/dataset2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)
	fmt.Println(df)

	dfSelect := df.Select([]string{"email", "gender"})
	fmt.Println(dfSelect)

	dfFiltered := df.Filter(
		dataframe.F{
			Colname:    "gender",
			Comparator: series.Eq,
			Comparando: "Male",
		},
	)

	fmt.Println(dfFiltered)

	// Filter and Select

	dfFilteredAndSelected := df.Filter(
		dataframe.F{
			Colname:    "age",
			Comparator: series.GreaterEq,
			Comparando: 18,
		},
	).Select([]string{"email"})

	fmt.Println(dfFilteredAndSelected)

}
