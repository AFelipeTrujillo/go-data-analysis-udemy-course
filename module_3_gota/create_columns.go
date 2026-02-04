package module3

import (
	"fmt"
	"log"

	"github.com/AFelipeTrujillo/go-data-analysis-udemy-course/utils"
	"github.com/go-gota/gota/series"
)

func CreateColumns() {

	df := utils.CreateDataFrame("./datasets/dataset2.csv")
	fmt.Println(df.Names())

	ages, err := df.Col("age").Int()
	if err != nil {
		log.Fatal(err)
	}

	doubleAges := make([]int, len(ages))

	for i, age := range ages {
		doubleAges[i] = age * 2
	}

	newColDoubleAge := series.New(doubleAges, series.Int, "double_ages")

	df = df.Mutate(newColDoubleAge)

	fmt.Println(df.Select([]string{"age", "double_ages"}))

}
