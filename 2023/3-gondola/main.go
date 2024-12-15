package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for _, line := range ReadFile("input.txt") {
		fmt.Println(line)
	}
}

func ReadFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
