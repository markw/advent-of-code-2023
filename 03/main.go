package main

import (
    "fmt"
    "strconv"
    msw "markw/lib"
)

func atoi(s string) int {
    n, err := strconv.Atoi(s)
    msw.Check(err)
    return n
}

type GridLoc struct {
    row int
    col int
    s string
}

type Grid struct {
    rows []string
}

func (grid Grid) charAt(r,c int) byte {
    if r < 0 || c < 0 || r > len(grid.rows) { return '.' }
    row := grid.rows[r]
    if c >= len(row) { return '.' }
    return row[c]
}

func isSymbol(ch byte) bool {
    return ch != '.' && !isDigit(ch)
}

func isDigit(ch byte) bool { return ch >= '0' && ch <= '9' }

func parseNums(s string, row int) []GridLoc {
    result := make([]GridLoc, 0)
    isNum := false
    numStr := ""
    startIndex := 0
    for n, ch := range s  {
        if isDigit(byte(ch)) {
            numStr = numStr + string(ch)
            if !isNum { 
                isNum = true
                startIndex = n
            }
        } else if isNum {
            isNum = false
            result = append(result, GridLoc{row, startIndex, numStr})
            numStr = ""
        }
    }
    // numbers at the end of the line
    if isNum {
        result = append(result, GridLoc{row, startIndex, numStr})
    }
    return result
}

func neighbors(loc GridLoc) []GridLoc {
    return []GridLoc {
        GridLoc{loc.row + 1, loc.col, loc.s}, // up
        GridLoc{loc.row - 1, loc.col, loc.s}, // down
        GridLoc{loc.row, loc.col - 1, loc.s}, // left
        GridLoc{loc.row, loc.col + 1, loc.s}, // right

        GridLoc{loc.row + 1, loc.col - 1, loc.s}, // up left
        GridLoc{loc.row + 1, loc.col + 1, loc.s}, // up right
        GridLoc{loc.row - 1, loc.col - 1, loc.s}, // down left
        GridLoc{loc.row - 1, loc.col + 1, loc.s}, // down right
    }
}

func isAdjacentToSymbol(grid Grid, loc GridLoc) bool {
    for i := 0; i < len(loc.s); i++ {
        curr := GridLoc{ loc.row, loc.col + i, loc.s }
        for _, neighbor := range neighbors(curr) {
            if isSymbol(grid.charAt(neighbor.row, neighbor.col)) {
                return true
            }
        }
    }
    return false
}

func buildInt(s string, n int) int {
    start, end := n, n
    for start >= 0 && isDigit(s[start]) {
        start--
        if start < 0 || !isDigit(s[start]) {
            start++
            break
        }
    }
    for end < len(s) && isDigit(s[end]) {
        end++
    }
    return atoi(s[start:end])
}

func main() {
    lines := msw.FileLines("input.txt")

    numberLocs := make([]GridLoc, 0)

    for n, line := range lines {
        numberLocs = append(numberLocs, parseNums(line, n)...)
    }

    grid := Grid{lines}

    isAdjacent := func(loc GridLoc) bool { 
        return isAdjacentToSymbol(grid, loc) 
    }

    addLoc := func(accum int, loc GridLoc) int { 
        return accum + atoi(loc.s) 
    }

    adjacents := msw.Filter(numberLocs, isAdjacent)
    total := msw.Reduce(adjacents, addLoc, 0)

    fmt.Printf("Part 1: %d\n", total)

    gearLocs := make([]GridLoc, 0)
    for r, line := range lines {
        for c, ch := range line {
            if ch == '*' {
                gearLocs = append(gearLocs, GridLoc{r,c,"*"})
            }
        }
    }

    numsAdjacentToGears := msw.Map(gearLocs, func (loc GridLoc) *msw.Set[int] {
        result := msw.NewSet[int]()
        for _, neighbor := range neighbors(loc) {
            char := grid.charAt(neighbor.row, neighbor.col)
            if isDigit(char) {
                line := lines[neighbor.row]
                num := buildInt(line, neighbor.col)
                result.Add(num)
                // fmt.Printf("digit: %s line: %s, col %d built %d\n", string(char), line, neighbor.col, num)
            }
        }
        return result
    })

    totalGearRatios := 0
    for _, gear := range numsAdjacentToGears {
        nums := gear.Items()
        if len(nums) == 2 {
            ratio := nums[0] * nums[1]
            totalGearRatios += ratio
        }
    }

    fmt.Printf("Part 2: %d\n", totalGearRatios)
}
