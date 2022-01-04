package main

import (
	"fmt"
	"github.com/adityapwr/protobuf/simple"
)

func main() {
	doSimple()
}

func doSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm)

	fmt.Println("The ID is:", sm.GetId())

	return &sm
}
