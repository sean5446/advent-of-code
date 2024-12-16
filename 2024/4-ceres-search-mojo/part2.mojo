

fn process_grid(grid: List[String]) -> Int:
    var height = len(grid)
    var width = len(grid[0])
    var sum = 0

    for x in range(0, width - 2):
        for y in range(0, width - 2):

            # diagonal \
            backs = grid[y][x]
            backs += grid[y+1][x+1]
            backs += grid[y+2][x+2]

            # diagonal /
            slash = grid[y][x+2]
            slash += grid[y+1][x+1]
            slash += grid[y+2][x]

            if (backs == "MAS" or backs == "SAM") and (slash == "MAS" or slash == "SAM"):
                sum += 1

    return sum


fn main():
    try:
        var lines1: String = \
""".M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
.........."""
        var lines = open("input.txt", "r").read()
        var ans = process_grid(lines.split("\n"))
        print(ans)
    except:
        print("Error reading or processing file")

# 1888
