package main

import ( 
    "fmt" 
    str "strings"
    "strconv"
    msw "markw/lib"
)

func atoi(s string) int {
    n, err := strconv.Atoi(s)
    msw.Check(err)
    return n
}

func isdigit(c int32) bool { return c >= '0' && c <= '9' }

func firstDigitPart1(s string) string {
    index := str.IndexFunc(s, isdigit)
    return string(s[index])
}

func lastDigitPart1(s string) string {
    index := str.LastIndexFunc(s, isdigit)
    return string(s[index])
}

func toNum_part1(s string) int {
    first := firstDigitPart1(s)
    last := lastDigitPart1(s)
    return atoi(first + last)
}

func minDigitPos(arr []DigitPos) DigitPos {
    if len(arr) == 0 { panic("empty array") }
    result := arr[0]
    for _, p := range arr {
        if p.index < result.index {
            result = p
        }
    }
    return result
}

func maxDigitPos(arr []DigitPos) DigitPos {
    if len(arr) == 0 { panic("empty array") }
    result := arr[0]
    for _, p := range arr {
        if p.index > result.index {
            result = p
        }
    }
    return result
}

type DigitPos struct {
    index int
    digit string
}

var replacements = [][]string {
    {"one", "1"},
    {"two", "2"},
    {"three", "3"},
    {"four", "4"},
    {"five", "5"},
    {"six", "6"},
    {"seven", "7"},
    {"eight", "8"},
    {"nine", "9"},
}

func firstDigitPart2(s string) DigitPos {
    var positions = make([]DigitPos,0)
    for _, r := range replacements {
        var x = str.Index(s, r[0])
        if x >= 0 {
            positions = append(positions, DigitPos{x, r[1]})
        }
        x = str.IndexFunc(s, isdigit)
        if x >= 0 {
            positions = append(positions, DigitPos{x, string(s[x])})
        }
    }
    return minDigitPos(positions)
}

func lastDigitPart2(s string) DigitPos {
    var positions = make([]DigitPos,0)
    for _, r := range replacements {
        var x = str.LastIndex(s, r[0])
        if x >= 0 {
            positions = append(positions, DigitPos{x, r[1]})
        }
        x = str.LastIndexFunc(s, isdigit)
        if x >= 0 {
            positions = append(positions, DigitPos{x, string(s[x])})
        }
    }
    return maxDigitPos(positions)
}

func toNum_part2(s string) int {
    first := firstDigitPart2(s).digit
    last := lastDigitPart2(s).digit
    numStr := first + last
    // fmt.Printf("toNum 2: %s -> %s\n", s, numStr)
    return atoi(numStr)
}

func sum(ns []int) int {
    add := func (a int, b int) int { return a + b }
    return msw.Reduce(ns, add, 0)
}

func main() {
    input := msw.Filter(msw.FileLines("input.txt"), msw.StrNotEmpty)
    nums_part1 := msw.Map(input, toNum_part1)
    nums_part2 := msw.Map(input, toNum_part2)

    fmt.Printf("Part 1: %d\n", sum(nums_part1))
    fmt.Printf("Part 2: %d\n", sum(nums_part2))
}
