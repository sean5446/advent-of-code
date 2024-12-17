package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func numWinning(winning []string, having []string) int {
	winners := make(map[string]bool)
	count := 0
	for _, num := range winning {
		if num == "" {
			continue
		}
		winners[num] = true
	}
	for _, num := range having {
		if winners[num] {
			count++
		}
	}
	return count
}

func part1(input string) int {
	// 	input = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	sum := 0

	for _, row := range strings.Split(input, "\n") {
		winning := strings.Split(strings.Split(strings.Split(row, ":")[1], "|")[0], " ")
		having := strings.Split(strings.Split(row, "|")[1], " ")
		sum += int(math.Pow(2, float64(numWinning(winning, having)-1)))
	}

	return sum
}

func part2(input string) int {
	// 	input = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	sum := 0
	numInput := len(strings.Split(input, "\n"))
	cardCounts := make([]int, numInput)

	for i := range numInput {
		cardCounts[i] = 1
	}

	for index, row := range strings.Split(input, "\n") {
		winning := strings.Split(strings.Split(strings.Split(row, ":")[1], "|")[0], " ")
		having := strings.Split(strings.Split(row, "|")[1], " ")
		numWinning := numWinning(winning, having)
		for i := 1; i <= numWinning; i++ {
			cardCounts[index+i] += cardCounts[index]
		}
	}

	for x := range numInput {
		// fmt.Printf("Card %d %d\n", x+1, cardCounts[x])
		sum += cardCounts[x]
	}

	return sum
}
