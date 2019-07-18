package main

import (
	"fmt"
)

func pointers() {
	fmt.Println("Pointers Print START:")

	x := 542
	fmt.Println("Value of x:", x, "Address of x:", x)
	b := &x
	fmt.Println("Value of b:", b)
	fmt.Printf("Type of b: %T\n ", b)
	fmt.Println("Pointers Print END")
}
