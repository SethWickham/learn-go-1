package main

import (
	"fmt"
)

//Here we are creating two structs in go
// a struct is basically a data structure that predefines our values
// this is a really great feature in go which helps us to build  our code
// while maintaining clarity
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

	//then we are creating our composite literal and filling in our Values
	// from our predefined fields above.
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
	// that it is of TYPE zooAnimalNames, instead we can just call our outer TYPE.
	fmt.Println("is", zNames.tigerName, "dangerous?", ZooNorCal.tigerDanger)

	fmt.Println("structs1 Print END")
}
