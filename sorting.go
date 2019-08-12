package main

import (
	"fmt"
	"sort"
)

func sorting() {
	fmt.Println("sorting print START")
	//this is our slice of int
	numbers := []int{100, 7, 4, 10}
	fmt.Println("unsorted numbers:", numbers)

	//the sort.Ints turns our unordered list (of TYPE slice of int) into an ordered list from least to greatest in value
	sort.Ints(numbers)
	fmt.Println("sorted numbers:", numbers)
	fmt.Printf("numbers TYPE: %T\n", numbers)

	// This is our slice of string
	cast := []string{"Becky", "Judy", "Tony", "DannyBoy"}
	fmt.Println("unsorted cast:", cast)

	// using the sort.Strings we are turning our unordered slice of string into an ordered
	//list where the names are arranged in Alphabetical order.
	sort.Strings(cast)

	fmt.Println("sorted cast", cast)
	fmt.Printf("cast TYPE: %T\n", cast)
	fmt.Println("sorting print END")
}
