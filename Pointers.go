package main

import (
	"fmt"
)

//here is our pointers func which contains our  examples of Pointers used in go
// pointers are helpful if you have a large piece of data and you don't want to pass that large piece of data around
// your program. Instead we can just pass an Address where that data is stored.
func pointers() {
	fmt.Println("Pointers Print START:")
	// using a declaration assignment to assign the VALUE of 542 to x
	// therefore x is of TYPE int
	x := 542
	fmt.Println("Value of x:", x, "Address of x:", x)

	// here we creating a Variable and assigning its value to the ADDRESS of x using the & operator
	// the & shows us where the ADDRESS is of the Value that was stored in the computer's memory.
	b := &x
	fmt.Println("Value of b:", b)
	// with this Print we can see that our TYPE is *int (pointer to an int)
	// which is a completely different TYPE then int
	fmt.Printf("Type of b: %T\n ", b)

	// then we are able to locate (“dereference”) our Value from our ADDRESS by using the * operator
	fmt.Println("Print Value stored at ADDRESS:", *b)
	fmt.Println("Pointers Print END")
}
