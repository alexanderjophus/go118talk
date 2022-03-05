package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

type id int

func main() {
	fmt.Println(min(42, 24))
	fmt.Println(min("abc", "def"))
	fmt.Println(min(4.2, 8.9))
	fmt.Println(min(id(42), id(24)))
}
