package main

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
)

func main() {

	df := dataframe.LoadStructs([]struct {
		Name      string
		Age       int
		CreatedAt string
	}{
		{"Alice", 30, "2022-01-01"},
		{"Bob", 25, "2022-02-15"},
	})

	fmt.Println(df)

}
