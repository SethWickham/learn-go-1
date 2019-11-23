package main

import (
	"encoding/json"
	"fmt"
)

// we use marshal when we are needing to convert our go code to the json format
//at the top we bring in our encoding/json package which allows us to run our code

// our fields must all be capitilized to be exported by marshal, even though our struct doesn't need to be
type greetings struct {
	Hello string
	Hi    string
	Howdy string
}

func marshal() {
	fmt.Println("marshal1 Print START")

	//composite literal
	h1 := greetings{
		Hello: "english",
		Hi:    "shortened version",
		Howdy: "accent version",
	}
	// we name this bs because we get returned a byte slice.
	bs, err := json.Marshal(h1)
	if err != nil {
		fmt.Println("error:", err)
	}
	// then we are taking our byte converting into a string and printing it.
	fmt.Println("string of json:", string(bs))

	fmt.Println("marshal1 Print END")

}
