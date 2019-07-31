package main

import (
	"encoding/json"
	"fmt"
)

// see json to go website as a tool that transforms json into go code
type jsonGreetings struct {
	Hello string `json:"Hello"`
	Hi    string `json:"Hi"`
	Howdy string `json:"Howdy"`
}

func unmarshal() {
	fmt.Println("unmarshall print START")

	jsonString := `[{"Hello":"english","Hi":"shortened version","Howdy":"accent version"}]`
	bs := []byte(jsonString)
	fmt.Printf("%T\n", jsonString)
	fmt.Printf("%T\n", bs)

	var greetings []jsonGreetings

	err := json.Unmarshal(bs, &greetings)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("unmarshaled json:", greetings)
	fmt.Println("unmarshall print END")

}
