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

func Sum(array []int) int {
	var result int = 0
	for _, value := range array {
		result += int(value)
	}
	return result
}
