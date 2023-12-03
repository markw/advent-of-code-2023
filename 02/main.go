package main

import (
    "fmt"
    "strconv"
    str "strings"
    msw "markw/lib"
)

type RGB struct {
    red int
    green int
    blue int
}

type Game struct {
    id int
    sets []RGB
    maxRed func() int
    maxBlue func() int
    maxGreen func() int
    power func() int
}

func isValidGame(g Game) bool {
    return g.maxRed() <= 12 && g.maxGreen() <= 13 && g.maxBlue() <= 14
}

func max(ns []int) int {
    if len(ns) == 0 { panic("empty array") }
    greaterOf := func(a,b int) int {
        if b > a { return b } 
        return a
    }
    return msw.Reduce(ns, greaterOf, ns[0])
}

func newGame(id int, sets []RGB) Game {
    var g Game

    reds := func(g Game) []int { return msw.Map(g.sets, func (rgb RGB) int {return rgb.red}) }
    blues := func(g Game) []int { return msw.Map(g.sets, func (rgb RGB) int {return rgb.blue}) }
    greens := func(g Game) []int { return msw.Map(g.sets, func (rgb RGB) int {return rgb.green}) }

    g.id = id
    g.sets = sets
    g.maxRed = func() int { return max(reds(g)) }
    g.maxBlue = func() int { return max(blues(g)) }
    g.maxGreen = func() int { return max(greens(g)) }
    g.power = func() int { return g.maxRed() * g.maxBlue() * g.maxGreen() }
    return g
}

func parseRGB(s string) RGB {
    var r,g,b int
    tokens := str.Split(s, ",")
    for _, token := range tokens {
        if str.Contains(token, "red") { fmt.Sscanf(token, "%d red", &r) }
        if str.Contains(token, "blue") { fmt.Sscanf(token, "%d blue", &b) }
        if str.Contains(token, "green") { fmt.Sscanf(token, "%d green", &g) }
    }
    return RGB{r,g,b}
}

func parseGame(s string) Game {
    var gameId int
    fmt.Sscanf(s, "Game %d:", &gameId)

    rgbSets := make([]RGB, 0)

    rhs := str.Trim(str.Split(s, ":")[1], " ")
    setStrs := str.Split(rhs, ";")
    for _, rgbStr := range setStrs {
        rgbStr := str.Trim(rgbStr, " ")
        rgbSets = append(rgbSets, parseRGB(rgbStr))
    }

    return newGame(gameId, rgbSets)
}

func main() {
    lines := msw.Filter(msw.FileLines("input.txt"), msw.StrNotEmpty)
    games := msw.Map(lines, parseGame)
    validGames := msw.Filter(games, isValidGame)
    addGameId := func(accum int, g Game) int { return accum + g.id }
    fmt.Println("Part 1: " + strconv.Itoa(msw.Reduce(validGames, addGameId, 0)))

    powers := msw.Map(games, func(g Game) int { return g.power() })
    plus := func(a, b int) int { return a + b }
    fmt.Println("Part 2: " + strconv.Itoa(msw.Reduce(powers, plus, 0)))
}
