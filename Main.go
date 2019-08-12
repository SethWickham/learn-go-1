package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from func main")

	// funcs from other files
	// the func names come from the code
	//inside each func, which contain
	// examples of go fundamentals
	pointers()
	structs1()
	marshal()
	unmarshal()
	sorting()
	sortCustom()
}
