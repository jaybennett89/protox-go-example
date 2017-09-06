package main

import (
	"fmt"
	"os"

	"github.com/jaybennett89/protoexperiment/experiment"
	"github.com/jaybennett89/protoexperiment/protox"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func test() error {
	// create an example with valid data and do a simple marshal/unmarshal test.
	example := &experiment.Example{
		Message:    "Hello world.",
		Percentage: 99,
		List:       []int32{1},
	}

	data, err := protox.Marshal(example)
	if err != nil {
		return err
	}

	copy := &experiment.Example{}
	err = protox.Unmarshal(data, copy)
	if err != nil {
		return err
	}

	fmt.Printf("example: %#v\n", example)
	fmt.Printf("copy: %#v\n", copy)

	return nil
}
