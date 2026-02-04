package module3

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func CombineFilters() {

	file, err := os.Open("./datasets/dataset2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)

	dfAnd := df.Filter(dataframe.F{
		Colname:    "gender",
		Comparator: series.Eq,
		Comparando: "Male",
	}).Filter(dataframe.F{
		Colname:    "age",
		Comparator: series.GreaterEq,
		Comparando: 45,
	})

	fmt.Println(dfAnd)
	fmt.Println("Total:", dfAnd.Nrow())

	// AND FilterAggregation

	dfAndAggregation := df.FilterAggregation(
		dataframe.And,
		dataframe.F{
			Colname:    "age",
			Comparator: series.Greater,
			Comparando: 40,
		},
		dataframe.F{
			Colname:    "gender",
			Comparator: series.Eq,
			Comparando: "Male",
		},
	)

	fmt.Println(dfAndAggregation)

	// OR Aggregation

	dfOrAggregation := df.FilterAggregation(
		dataframe.Or,
		dataframe.F{
			Colname:    "gender",
			Comparator: series.Eq,
			Comparando: "Male",
		},
		dataframe.F{
			Colname:    "age",
			Comparator: series.Greater,
			Comparando: 40,
		},
	)

	fmt.Println(dfOrAggregation)

	// NOT IN FILTER

	dfNotIn := df.Filter(dataframe.F{
		Colname:    "email",
		Comparator: series.In,
		Comparando: []string{"kbagnold6@un.org"},
	})

	fmt.Println(dfNotIn)
}
