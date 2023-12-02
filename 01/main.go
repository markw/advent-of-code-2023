package main

import ( 
    "fmt" 
    "os" 
    str "strings"
    "strconv"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func Map[T, V any](ts []T, fn func(T) V) []V {
    result := make([]V, len(ts))
    for i, t := range ts {
        result[i] = fn(t)
    }
    return result
}

func Filter[T any](ts []T, f func(T) bool) []T {
    result := make([]T, 0)
    for _, t := range ts {
        if f(t) {
            result = append(result, t)
        }
    }
    return result
}

func filterStr(s string, f func(int32) bool) string {
    return string(Filter([]rune(s), f))
}

func atoi(s string) int {
    n, err := strconv.Atoi(s)
    check(err)
    return n
}

func isdigit(c int32) bool { return c >= '0' && c <= '9' }

func firstDigitIndex(s string) int {
    return str.IndexFunc(s, isdigit)
}

func lastDigitIndex(s string) int {
    return str.LastIndexFunc(s, isdigit)
}

func toNum_part1(s string) int {
    x := firstDigitIndex(s)
    y := lastDigitIndex(s)
    return atoi("" + string(s[x]) + string(s[y]))
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

func firstDigitWord(s string) (int, string) {
    var index = -1
    var rindex = -1

    for n, r := range replacements {
        var x = str.Index(s, r[0])
        if x >= 0 && (index < 0 || x < index) {
            index = x
            rindex = n
        }
    }
    if index >= 0 {
        return index, replacements[rindex][1]
    }
    return -1, ""
}

func lastDigitWord(s string) (int, string) {
    var index = -1
    var rindex = -1

    for n, r := range replacements {
        var x = str.LastIndex(s, r[0])
        if x > index {
            index = x
            rindex = n
        }
    }
    if index >= 0 {
        return index, replacements[rindex][1]
    }
    return -1, ""
}

func toNum_part2(s string) int {
    var first, last string
    x := firstDigitIndex(s)
    x0, rx := firstDigitWord(s)

    if rx != "" && x0 < x {
        first = rx 
    } else {
        first = string(s[x])
    }

    y := lastDigitIndex(s)
    y0, ry := lastDigitWord(s)
    if ry != "" && y0 > y{
        last = ry 
    } else {
        last = string(s[y])
    }

    result := "" + first + last
    // fmt.Printf("toNum 2: %s -> %s\n", s, result)

    return atoi(result)
}

func fileLines(fileName string) []string {
    data, err := os.ReadFile(fileName)
    check(err)
    return str.Split(string(data), "\n")
}

func strNotEmpty(s string) bool { return len(s) > 0 }

func sum(ns []int) int {
    total := 0
    for _, n := range ns {
        total += n
    }
    return total
}

func main() {
    input := Filter(fileLines("input.txt"), strNotEmpty)
    nums_part1 := Map(input, toNum_part1)
    nums_part2 := Map(input, toNum_part2)

    fmt.Printf("Part 1: %d\n", sum(nums_part1))
    fmt.Printf("Part 2: %d\n", sum(nums_part2))
}
