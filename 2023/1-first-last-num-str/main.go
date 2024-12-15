package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := ReadFile("input.txt")
	sum := 0

	for _, line := range lines {
		fmt.Println(line)

		// debug: What do you get if the input is twone, or 1threeight?
		if len(line) == 0 {
			break
		}

		// re := regexp.MustCompile(`(\d)`)
		// matches := re.FindAllStringSubmatch(lines, -1)
		// first := matches[0][0]
		// last := matches[len(matches)-1][0]

		forward := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine|ten)`)
		matches := forward.FindAllStringSubmatch(line, -1)
		first := stringToInt(matches[0][0])

		backward := regexp.MustCompile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)
		matches = backward.FindAllStringSubmatch(reverse(line), -1)
		last := stringToInt(reverse(matches[0][0]))

		combo, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Printf("input: %s %s\n", first, last)
			panic(err)
		}
		sum += combo

		fmt.Println(first, last)
	}
	fmt.Println(sum)
}

func reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func stringToInt(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return number
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
