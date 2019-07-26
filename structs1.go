package main

import (
	"fmt"
)

type zooAnimals struct {
	elephantName string
	toucanName   string
	tigerName    string
}

func structs1() {
	zooCA := zooAnimals{
		elephantName: "Sammy",
		toucanName:   "Tory",
		tigerName:    "George",
	}

	fmt.Println("structs1 Print START")
	fmt.Println("The Elephant:", zooCA.elephantName)
	fmt.Println("The Toucan:", zooCA.toucanName)
	fmt.Println("The Tiger:", zooCA.tigerName)

	fmt.Println("structs1 Print END")
}
