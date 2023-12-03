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

type NumLoc struct {
    row int
    col int
    s string
}

type Grid struct {
    rows []string
    charAt func(r,c int) rune
}

func isSymbol(ch int32) bool {
    return ch != '.' && !isDigit(ch)
}

func isDigit(ch int32) bool { return ch >= '0' && ch <= '9' }

func parseNums(s string, row int) []NumLoc {
    result := make([]NumLoc, 0)
    isNum := false
    numStr := ""
    startIndex := 0
    for n, ch := range s  {
        if isDigit(ch) {
            numStr = numStr + string(ch)
            if !isNum { 
                isNum = true
                startIndex = n
            }
        } else if isNum {
            isNum = false
            result = append(result, NumLoc{row, startIndex, numStr})
            numStr = ""
        }
    }
    // numbers at the end of the line
    if isNum {
        result = append(result, NumLoc{row, startIndex, numStr})
    }
    return result
}

func neighbors(loc NumLoc) []NumLoc {
    return []NumLoc {
        NumLoc{loc.row + 1, loc.col, loc.s}, // up
        NumLoc{loc.row - 1, loc.col, loc.s}, // down
        NumLoc{loc.row, loc.col - 1, loc.s}, // left
        NumLoc{loc.row, loc.col + 1, loc.s}, // right

        NumLoc{loc.row + 1, loc.col - 1, loc.s}, // up left
        NumLoc{loc.row + 1, loc.col + 1, loc.s}, // up right
        NumLoc{loc.row - 1, loc.col - 1, loc.s}, // down left
        NumLoc{loc.row - 1, loc.col + 1, loc.s}, // down right
    }
}

func isAdjacentToSymbol(grid Grid, loc NumLoc) bool {
    for i := 0; i < len(loc.s); i++ {
        curr := NumLoc{ loc.row, loc.col + i, loc.s }
        for _, neighbor := range neighbors(curr) {
            if isSymbol(grid.charAt(neighbor.row, neighbor.col)) {
                return true
            }
        }
    }
    return false
}

func newGrid(lines []string) Grid {
    var grid Grid
    grid.rows = lines
    grid.charAt = func (r,c int) rune { 
        if r < 0 || c < 0 || r > len(grid.rows) { return '.' }
        row := grid.rows[r]
        if c >= len(row) { return '.' }
        return rune(row[c]) 
    } 
    return grid
}

func main() {
    lines := msw.FileLines("input.txt")

    numLocs := make([]NumLoc, 0)

    for n, line := range lines {
        numLocs = append(numLocs, parseNums(line, n)...)
    }

    grid := newGrid(lines)

    isAdjacent := func(loc NumLoc) bool { 
        return isAdjacentToSymbol(grid, loc) 
    }

    addLoc := func(accum int, loc NumLoc) int { 
        return accum + atoi(loc.s) 
    }

    adjacents := msw.Filter(numLocs, isAdjacent)
    total := msw.Reduce(adjacents, addLoc, 0)

    fmt.Printf("Part 1: %d\n", total)

}
