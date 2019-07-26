package main

import (
	"fmt"
)

type zooAnimalNames struct {
	elephantName string
	toucanName   string
	tigerName    string
}

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

	fmt.Println("is", zNames.tigerName, "dangerous?", ZooNorCal.tigerDanger)

	fmt.Println("structs1 Print END")
}
