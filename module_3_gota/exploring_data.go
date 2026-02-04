package module3

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func ExploringData() {

	file, err := os.Open("./datasets/dataset2.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)
	fmt.Println(df.Subset(RangeInts(0, 5)))
	fmt.Println(df.Subset([]int{0, 1, 2}))
	n := df.Nrow()
	fmt.Println(df.Subset(RangeInts(n-5, n)))

	// Print a column
	fmt.Println("Print one colum")
	fmt.Println(df.Col("age"))

	// Describe
	fmt.Println(df.Describe())

}

func RangeInts(from, to int) []int {
	size := to - from
	index := make([]int, size)

	for i := 0; i < size; i++ {
		index[i] = from + i
	}

	return index

}
