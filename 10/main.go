package main

import (
    "fmt"
    msw "markw/lib"
)

type Addr struct {
    row, col int
}

func (a Addr) north() Addr { return Addr{a.row - 1, a.col} }
func (a Addr) south() Addr { return Addr{a.row + 1, a.col} }
func (a Addr) east()  Addr { return Addr{a.row, a.col + 1} }
func (a Addr) west()  Addr { return Addr{a.row, a.col - 1} }

type Cell struct {
    a0, a1 Addr
}

type Grid [140][140]Cell

func (g Grid) get(a Addr) Cell { return g[a.row][a.col] }

func parseCell(row, col int, ch rune) Cell {
    addr := Addr{row,col}
    noAddr := Addr{-1,-1}
    empty := Cell{noAddr, noAddr}
    switch ch {
        case '|': return Cell{addr.north(), addr.south()}
        case '-': return Cell{addr.east(),  addr.west()}
        case 'L': return Cell{addr.north(), addr.east()}
        case 'J': return Cell{addr.north(), addr.west()}
        case '7': return Cell{addr.south(), addr.west()}
        case 'F': return Cell{addr.south(), addr.east()}
        case '.': return empty
        case 'S': return empty
    }
    return empty
}

func nextAddr(a Addr, c Cell) Addr {
    // fmt.Printf("nextAddr a %v c %v\n", a, c)
    if (c.a0 == a) { return c.a1 }
    if (c.a1 == a) { return c.a0 }
    panic(fmt.Sprintf("nextAddr a %v c %v\n", a, c))
}

func main () {
    grid := new(Grid)
    var startAddr Addr
    lines := msw.FileLines("input.txt")
    for row, line := range lines {
        for col, ch := range line {
            if ch == 'S' {
                // fmt.Printf("S at %d,%d\n", row, col)
                startAddr = Addr{row,col}
            }
            grid[row][col] = parseCell(row, col, ch)
        }
    }

    curr := startAddr
    next := startAddr.south()
    steps := 1
    for next != startAddr {
        tmp := nextAddr(curr, grid.get(next))
        curr = next
        next = tmp
        steps++
    }
    fmt.Printf("Part 1: %d\n", steps/2)
}

