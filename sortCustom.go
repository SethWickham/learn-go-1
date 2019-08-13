package main

import (
	"fmt"
	"sort"
)

// FireTruck : this is the basic data structure we are working with in this file
type FireTruck struct {
	NameOfDriver string
	TruckID      int
}

// This was pulled from godoc.org and slightly altered to fit what was needed.
// This file contains code that allows us to sort through data that has a custom type
// which we have created

// Here we are implementing the sort.Interface for the []FireTruck based on the TruckID field.

// ByTruckID :  allows us to order the trucks by ID
type ByTruckID []FireTruck

func (id ByTruckID) Len() int           { return len(id) }
func (id ByTruckID) Swap(i, j int)      { id[i], id[j] = id[j], id[i] }
func (id ByTruckID) Less(i, j int) bool { return id[i].TruckID < id[j].TruckID }

func sortCustom() {
	fmt.Println("sortCustom print START")

	trucks := []FireTruck{
		{"Fred", 19756},
		{"Jared", 18894},
		{"Dennis", 23735},
		{"Lucy", 18459},
		{"Mark", 29876}}

	fmt.Println("unordered trucks:", trucks)

	// using the Sort method from our imported "sort" package will allow us call ByTruckID
	// then we are using conversion (ByTruckID(trucks)) to convert trucks to TYPE ByTruckID
	//so that any values from trucks now has access to our ByTruckID methods.
	sort.Sort(ByTruckID(trucks))
	fmt.Println("ordered trucks:", trucks)
	//now our trucks will be organized from the ID with the lowest number to the highest number
	fmt.Println("sortCustom print END")
}
