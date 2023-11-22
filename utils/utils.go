package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func ReadInput() ([]string, error) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	input := []string{}

	for _, line := range strings.Split(string(data), "\n") {
		input = append(input, line)
	}

	return input, err
}

func GetInput() error {
	// Check if input file already exists
	if _, err := os.Stat("input.txt"); !os.IsNotExist(err) {
		fmt.Println("Input file already exists, skipping GetInput")
		return errors.New("Input file already exists")
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Read Session ID
	sessionID, err := os.ReadFile("../.sessionId")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Create a .sessionId in the current year folder with cookie from authenticated session.")
		return err
	}

	// Determine advent date
	path := strings.Split(cwd, "/")
	year := path[len(path)-2]
	day := strings.Split(path[len(path)-1], "day")[1]

	// Get input from AoC
	client := http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"Cookie": []string{fmt.Sprintf("session=%s", string(sessionID))},
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// Check if session cookie is valid
	body, err := io.ReadAll(resp.Body)
	if strings.Contains(string(body), "Puzzle inputs differ by user") {
		fmt.Println("You need to provide a valid session cookie in order to download the input.")
		return err
	}

	// Create dummy input file
	out, err := os.Create("input.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer out.Close()

	// Write input to file
	_, err = out.WriteString(string(body))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func Sum(array []int) int {
	var result int = 0
	for _, value := range array {
		result += int(value)
	}
	return result
}

func ReverseBytes(bytes []byte) []byte {
	for i := 0; i < len(bytes)/2; i++ {
		j := len(bytes) - i - 1
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}

func ReverseInts(ints []int) []int {
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
	return ints
}
