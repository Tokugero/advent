package main

import (
	"fmt"
	"io/ioutil"
)

func ReadInput() (string, error) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(data), err
}
