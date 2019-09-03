package main

import "fmt"

func main() {
	multidimensionalSlice := [][]string{{"James", "Bond", "Shaken not stirred"}, {"Miss", "Moneypenny", "Hello James"}}
	fmt.Println(multidimensionalSlice)

	for _, sliceOfslice := range multidimensionalSlice {
		for _, itemInSlice := range sliceOfslice {
			fmt.Println(itemInSlice)
		}
	}
}
