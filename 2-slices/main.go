package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	// sorting
	fmt.Println("SORTING")
	myList := []int{3, 2, 1}
	slices.Sort(myList)
	fmt.Println(myList)
	slices.SortFunc(myList,
		func(a, b int) bool {
			return a > b
		},
	)
	fmt.Println(myList)

	fmt.Println("IS SORTED")
	fmt.Println(slices.IsSorted([]int{1, 2, 3}))
	fmt.Println(slices.IsSorted([]int{3, 2, 1}))
	fmt.Println(slices.IsSortedFunc([]int{3, 2, 1},
		func(a, b int) bool {
			return a > b
		},
	))

	// manipulating slices
	fmt.Println("MANIPULATING SLICES")
	myList = []int{1, 1, 1, 1, 2, 3}
	// slices.Compact(myList) // gotcha (must be `myList = ...`)
	myList = slices.Compact(myList)
	fmt.Println(myList)
}
