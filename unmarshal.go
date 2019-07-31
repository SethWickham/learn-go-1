package main

import (
	"encoding/json"
	"fmt"
)

//I used the json to go website in order to create this struct
type jsonGreetings struct {
	Hello string `json:"Hello"`
	Hi    string `json:"Hi"`
	Howdy string `json:"Howdy"`
}

func unmarshal() {
	fmt.Println("unmarshall print START")

	// here we place our json code inside of backticks and it's then read as a string
	jsonString := `[{"Hello":"english","Hi":"shortened version","Howdy":"accent version"}]`

	//here we are converting our string of json into a bite slice
	// becuase it is required by the encoding/json package to work correctl
	bs := []byte(jsonString)

	// our type of bs will print []uint8 which is slice of byte
	fmt.Printf("%T\n", bs)

	var greetings []jsonGreetings

	// this information is found at golang.org.
	//Here we are running our code that will Unmarshal our json.
	// It does this by taking in our slice of byte which contains our json
	// and the address (ampersand points to the location) of our jsonGreetings data structure
	error := json.Unmarshal(bs, &greetings)
	//if error is not equal to nil then we print the error so we can see what went wrong
	if error != nil {
		fmt.Println(error)
	}
	//finally our print which is our unmarshaled json YAY!
	fmt.Println("unmarshaled json:", greetings)
	fmt.Println("unmarshall print END")

}
