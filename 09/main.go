package main

import (
    "fmt"
    str "strings"
    msw "markw/lib"
)

func allZeroes(ns []int) bool {
    for _, n := range ns {
        if n != 0 { return false }
    }
    return true
}

func diffsForSeq(ns []int) []int {
    result := make([]int, len(ns) - 1)
    for i := 0; i < len(result); i++ {
        result[i] = ns[i+1] - ns[i]
    }
    return result
}

func buildDiffs(ns []int) [][]int {
    diffSeqs := make([][]int, 1)
    diffSeqs[0] = ns
    for !allZeroes(diffSeqs[0]) {
        nextDiffs := diffsForSeq(diffSeqs[0])
        diffSeqs = append([][]int{nextDiffs}, diffSeqs...)
    }
    return diffSeqs
}

func extrapolateForward(ns [][]int) int {
    result := 0
    for _, r := range ns {
        result += r[len(r)-1]
    }
    return result
}

func extrapolateBackward(ns [][]int) int {
    result := 0
    for i := 0; i < len(ns); i++ {
        n := ns[i][0]
        result = n - result
    }
    return result
}

func main () {
    lines := msw.Filter(msw.FileLines("input.txt"), msw.StrNotEmpty)
    part1, part2 := 0, 0
    for _, line := range lines {
        nums := msw.Map(str.Fields(line), msw.ParseInt)
        diffs := buildDiffs(nums)
        part1 += extrapolateForward(diffs)
        part2 += extrapolateBackward(diffs)
    }
    fmt.Printf("Part 1: %d\n", part1)
    fmt.Printf("Part 2: %d\n", part2)
}
