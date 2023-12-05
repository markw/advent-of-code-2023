package main

import ( 
    "fmt" 
    "strconv"
    msw "markw/lib"
    "reflect"
)

func atoi(s string) int {
    n, err := strconv.Atoi(s)
    msw.Check(err)
    return n
}

func isdigit(c int32) bool { return c >= '0' && c <= '9' }

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

func sum(ns []int) int {
    add := func (a int, b int) int { return a + b }
    return msw.Reduce(ns, add, 0)
}

func startsWith(s string, target string) bool {
    if len(target) > len(s) { return false }
    rune0 := []rune(s[0:len(target)])
    rune1 := []rune(target)
    // fmt.Printf("s %s target %s\n", s, target)
    return reflect.DeepEqual(rune0, rune1)
}

func extractIntPart1(s string) int {
    digits := make([]string,0)
    for i := 0; i < len(s); i++ {
        if isdigit(rune(s[i])) {
            digits = append(digits, string(s[i]))
        }
    }
    numStr := digits[0] + digits[len(digits)-1]
    return atoi(numStr)
}

func extractIntPart2(s string) int {
    digits := make([]string,0)
    for i := 0; i < len(s); i++ {
        if isdigit(rune(s[i])) {
            digits = append(digits, string(s[i]))
        }
        substr := s[i:]
        for _, r:= range replacements {
            if startsWith(substr, r[0]) {
                digits = append(digits, r[1])
            }
        }
    }
    numStr := digits[0] + digits[len(digits)-1]
    return atoi(numStr)
}

func main() {
    input := msw.Filter(msw.FileLines("input.txt"), msw.StrNotEmpty)
    numsPart1 := msw.Map(input, extractIntPart1)
    numsPart2 := msw.Map(input, extractIntPart2)

    fmt.Printf("Part 1: %d\n", sum(numsPart1))
    fmt.Printf("Part 2: %d\n", sum(numsPart2))
}
