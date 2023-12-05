package main

import (
    "fmt"
    str "strings"
    msw "markw/lib"
)

type Range struct {
    dest int
    src int
    length int
}

type Alamanac struct {
    seeds []int
    maps [][]Range
}

func parseAlamanac(ss []string) Alamanac {
    var seeds []int
    var maps [][]Range

    for _,s := range ss {
        if str.Index(s, "seeds:") == 0 {
            tokens := str.Fields(s)[1:]
            seeds = msw.Map(tokens, msw.ParseInt)
        } else if str.Index(s, "map:") > 1 {
            maps = append(maps, make([]Range, 0))
        } else if len(s) > 0 {
            ns := msw.Map(str.Fields(s), msw.ParseInt)
            maps[len(maps)-1] = append(maps[len(maps)-1], Range{ns[0], ns[1], ns[2]})
        }
    }
    return Alamanac{seeds, maps}
}

func convert (n int, r Range) int {
    s0, s1 := r.src, r.src + r.length - 1
    if s0 <= n && n <= s1 {
        return n + r.dest - r.src
    }
    return n
}

func lookup(n int, rs []Range) int {
    for _, r := range rs {
        c := convert(n, r)
        if (c != n) {
            return c
        }
    }
    return n
}

func processSeed(seed int, almanac Alamanac) int {
    var curr int
    curr = seed
    for _, r := range almanac.maps {
        curr = lookup(curr, r)
    }
    return curr
}

func minLocationPart1(almanac Alamanac) int {
    locations := make([]int, len(almanac.seeds))
    for i, seed := range almanac.seeds {
        locations[i] = processSeed(seed, almanac)
    }
    return msw.Min(locations)
}

func minLocationPart2(almanac Alamanac) int {
    minLocation := msw.Max(almanac.seeds) // should be good enough
    for i := 0; i < len(almanac.seeds); i += 2 {
        seed := almanac.seeds[i]
        length := almanac.seeds[i+1]
        for j := 0; j < length; j++ {
            location := processSeed(seed, almanac)
            if location < minLocation { 
                minLocation = location
            }
            seed++
        }
    }
    return minLocation
}


func main() {
    lines := msw.FileLines("input.txt")
    almanac := parseAlamanac(lines)
    fmt.Printf("Part 1: %d\n", minLocationPart1(almanac))
    fmt.Printf("Part 2: %d\n", minLocationPart2(almanac))
}
