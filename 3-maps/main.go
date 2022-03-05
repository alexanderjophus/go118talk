package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	fmt.Println(maps.Keys(myMap))
	fmt.Println(maps.Values(myMap))
}
