package main

import (
	"fmt"
)

//Here we are creating two structs in go
type zooAnimalNames struct {
	elephantName string
	toucanName   string
	tigerName    string
}

//  here we are embedding our zooAnimalNames struct
// into in our dangerousZooAnimals struct so that the Values can be accessed later on
type dangerousZooAnimals struct {
	zooAnimalNames
	elephantDanger bool
	toucanDanger   bool
	tigerDanger    bool
}

func structs1() {

	ZooNorCal := dangerousZooAnimals{
		zooAnimalNames: zooAnimalNames{
			elephantName: "Sammy",
			toucanName:   "Tory",
			tigerName:    "George",
		},

		elephantDanger: false,
		toucanDanger:   false,
		tigerDanger:    true,
	}

	zNames := zooAnimalNames{
		elephantName: "Sammy",
		toucanName:   "Tory",
		tigerName:    "George",
	}

	fmt.Println("structs1 Print START")
	fmt.Println("The Elephant's Name is :", zNames.elephantName)
	fmt.Println("The Toucan's Name is :", zNames.toucanName)
	// in Go our inner TYPE gets promoted to the outer TYPE so when we want to call tigerDanger we don't need to specify
	// that it is of TYPE zooAnimals, instead we can just call our TYPE from the outer TYPE zooNorCal.
	fmt.Println("is", zNames.tigerName, "dangerous?", ZooNorCal.tigerDanger)

	fmt.Println("structs1 Print END")
}
