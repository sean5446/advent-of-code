package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
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

var numReg = regexp.MustCompile("[0-9]")

var diffs = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func part1(input string) int {
	matrix := [][]string{}

	// collect special characters
	specialChars := map[string]bool{}
	for _, row := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(row, ""))
		for _, val := range strings.Split(row, "") {
			if !numReg.MatchString(val) && val != "." {
				specialChars[val] = true
			}
		}
	}

	specialCharsString := ""
	for k := range specialChars {
		specialCharsString += k
	}

	var sum int

	seen := map[[2]int]bool{}

	for r, row := range matrix {
		for c, val := range row {
			coords := [2]int{r, c}
			if seen[coords] {
				continue
			}
			seen[coords] = true

			// if we hit a number, collect the entire number and check along the way if it's
			// adjacent to a special char
			if numReg.MatchString(val) {
				hasAdjacentSpecialChar := false
				numStr := ""

				for j := 0; j+c < len(matrix[0]); j++ {
					char := row[c+j]
					// breaks on period or special character, loop itself breaks on out of range
					if !numReg.MatchString(char) {
						break
					}
					// keep collecting number
					numStr += char

					// check all 8 directions for special char
					for _, d := range diffs {
						dr, dc := r+d[0], c+j+d[1]
						if dr >= 0 && dr < len(matrix) && dc >= 0 && dc < len(matrix[0]) {
							if strings.ContainsAny(matrix[dr][dc], specialCharsString) {
								hasAdjacentSpecialChar = true
							}
							seen[[2]int{r, c + j}] = true
						}
					}
				}

				if hasAdjacentSpecialChar {
					var val, _ = strconv.Atoi(numStr)
					sum += val
				}

			}
		}
	}

	return sum
}

// getNumber returns -1 if a number is "not found" which could include the number
// already being seen
func getNumber(matrix [][]string, coord [2]int, seen map[[2]int]bool) int {
	if !numReg.MatchString(matrix[coord[0]][coord[1]]) {
		return -1
	}
	
	if seen[coord] {
		return -1
	}

	// go to the left most digit
	r, c := coord[0], coord[1]
	for c-1 >= 0 {
		if numReg.MatchString(matrix[r][c-1]) {
			c--
		} else {
			break
		}
	}

	numStr := ""

	for c < len(matrix[0]) && numReg.MatchString(matrix[r][c]) {
		numStr += matrix[r][c]
		seen[[2]int{r, c}] = true
		c++
	}

	var val, _ = strconv.Atoi(numStr)

	return val
}


func part2(input string) int {
	seen := map[[2]int]bool{}
	matrix := [][]string{}
	sum := 0

	for _, row := range strings.Split(input, "\n") {
		matrix = append(matrix, strings.Split(row, ""))
	}

	for r, rows := range matrix {
		for c, val := range rows {
			if val == "*" {
				nums := []int{}
				for _, diff := range diffs {
					dr, dc := r+diff[0], c+diff[1]
					if dr >= 0 && dr < len(matrix) && dc >= 0 && dc < len(matrix[0]) {
						foundNum := getNumber(matrix, [2]int{dr, dc}, seen)
						if foundNum != -1 {
							nums = append(nums, foundNum)
						}
					}
				}

				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}
	return sum
}
