package module3

import (
	"fmt"
	"strconv"

	"github.com/AFelipeTrujillo/go-data-analysis-udemy-course/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func WorkingWithNaN() {

	df := utils.CreateDataFrame("datasets/dataset_with_nan.csv")

	// type
	fmt.Println(df.Types())

	// filter nan
	dfWithOutNaN := df.Filter(dataframe.F{
		Colname:    "age",
		Comparator: series.CompFunc,
		Comparando: func(el series.Element) bool {
			return !el.IsNA()
		},
	})

	fmt.Println(dfWithOutNaN)

	ages := df.Col("age").Records()
	fmt.Println(ages)

	agesCleanList := make([]float64, len(ages))

	for i, valStr := range ages {
		value, err := strconv.ParseFloat(valStr, 10)
		if err != nil {
			value = 0
		}
		agesCleanList[i] = value
	}

	fmt.Println(agesCleanList)

	seriesAge := series.New(agesCleanList, series.Float, "Age Clean")

	fmt.Println(seriesAge)

	df = df.Mutate(seriesAge)

	fmt.Println(df)
}
