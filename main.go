package main

import (
	"fmt"
	"os"

	"github.com/jaybennett89/protoexperiment/example"
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
	e := &example.Example{
		Message:    "Hello world.",
		Percentage: 99,
		List:       []int32{1},
	}

	data, err := protox.Marshal(e)
	if err != nil {
		return err
	}

	copy := &example.Example{}
	err = protox.Unmarshal(data, copy)
	if err != nil {
		return err
	}

	fmt.Printf("example: %#v\n", e)
	fmt.Printf("copy: %#v\n", copy)

	return nil
}
