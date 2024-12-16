
# mojo doesn't have enumerate?
# for some reason col in ["XMAS", "SAMX"] doesn't work?
# accessing index on string with length 0 causes segfault
# can't get the type of variables??


fn process_grid(grid: List[String]) -> Int:
    var height = len(grid)
    var width = len(grid[0])
    var sum = 0

    # horizontal
    for x in range(0, width - 3):
        for y in range(0, height):
            row = grid[y][x:x+4]
            if row == "XMAS" or row == "SAMX":
                sum += 1

    # vertical
    for y in range(0, height - 3):
        for x in range(0, width):
            col = grid[y][x]
            col += grid[y+1][x]
            col += grid[y+2][x]
            col += grid[y+3][x]
            if col == "XMAS" or col == "SAMX":
                sum += 1

    # diagonal \
    for y in range(0, height - 3):
        for x in range(0, width - 3):
            diag = grid[y][x]
            diag += grid[y+1][x+1]
            diag += grid[y+2][x+2]
            diag += grid[y+3][x+3]
            if diag == "XMAS" or diag == "SAMX":
                sum += 1

    # diagonal /
    for y in range(0, height - 3):
        for x in range(3, width):
            diag = grid[y][x]
            diag += grid[y+1][x-1]
            diag += grid[y+2][x-2]
            diag += grid[y+3][x-3]
            if diag == "XMAS" or diag == "SAMX":
                sum += 1

    return sum


fn main():
    try:
        var lines1: String = \
"""MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"""
        var lines = open("input.txt", "r").read()
        var ans = process_grid(lines.split("\n"))
        print(ans)
    except:
        print("Error reading or processing file")


# 2514
