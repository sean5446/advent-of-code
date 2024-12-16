package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadFile("input.txt")
	sum := 0

	// for part 1
	max_red := 12
	max_green := 13
	max_blue := 14

	for _, line := range lines {
		sum += parseGame(line, max_red, max_green, max_blue)
	}
	fmt.Println(sum)
}

func parseGame(line string, max_red int, max_green int, max_blue int) int {
	// fmt.Println(line)

	// get game index
	re := regexp.MustCompile(`Game (\d+): `)
	matches := re.FindAllStringSubmatch(line, -1)
	idx := matches[0][1]
	index, err := strconv.Atoi(idx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(index)

	game := strings.Split(line, ": ")[1]
	draws := strings.Split(game, "; ")

	min_red := 0
	min_green := 0
	min_blue := 0

	for _, draw := range draws {
		// fmt.Println(draw)

		for _, numberColorPair := range strings.Split(draw, ", ") {
			// fmt.Println(combo)

			re := regexp.MustCompile(`(?P<Number>\d+) (?P<Color>\D+)`)
			matches := re.FindAllStringSubmatch(numberColorPair, -1)
			// fmt.Println(matches[0])
			num := matches[0][re.SubexpIndex("Number")]
			color := matches[0][re.SubexpIndex("Color")]

			number, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
			}

			if color == "red" {
				if number > min_red {
					min_red = number
				}
			} else if color == "green" {
				if number > min_green {
					min_green = number
				}
			} else if color == "blue" {
				if number > min_blue {
					min_blue = number
				}
			}
		}
	}

	// part 1: 2061
	// if min_red > max_red || min_green > max_green || min_blue > max_blue {
	// 	return 0
	// }
	// return index

	// part 2: 72596
	return min_red * min_green * min_blue
}
