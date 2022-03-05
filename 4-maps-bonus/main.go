package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type numeric interface {
	constraints.Integer | constraints.Float
}

func runFunc[M ~map[K]V, K comparable, V numeric](m M, fn func(V) V) M {
	for k, v := range m {
		m[k] = fn(v)
	}
	return m
}

func double[V numeric](v V) V {
	if v < 2 {
		var r V
		return r
	}
	return v * 2
}

func main() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	fmt.Println(m1)

	fmt.Println(runFunc(m1, double[int]))

	m2 := map[int]float64{
		1: 0.1,
		2: 0.2,
		3: 0.3,
		4: 0.4,
		5: 0.5,
	}

	fmt.Println(m2)

	fmt.Println(runFunc(m2, math.Sin))
}
