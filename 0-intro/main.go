package main

import "fmt"

type Account int

func CreateAccount(initValue int) Account {
	return Account(initValue)
}

func CreateAccountV2(initValue *int) Account {
	if initValue == nil {
		panic("unexpected initValue")
	}
	return Account(*initValue)
}

func main() {
	fmt.Println(CreateAccount(50))

	// cannot do this, need to create a variable
	// and point to it
	// fmt.Println(CreateAccountV2(&50))

	initValue := 50
	fmt.Println(CreateAccountV2(&initValue))

	fmt.Println(CreateAccountV2(intPtr(50)))
	fmt.Println(CreateAccountV2(ptr(50)))
}

func intPtr(i int) *int {
	return &i
}

func ptr[T any](v T) *T {
	return &v
}
