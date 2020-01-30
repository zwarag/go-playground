package main

import "fmt"

type A struct {
	a string
}

type B struct {
	A
}

func main() {

	a := &A{
		a: "yo",
	}

	b := &B{
		A{
			a: "swag",
		},
	}

	fmt.Println(a)
	fmt.Println(b)
}
