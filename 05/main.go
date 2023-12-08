package main

import (
    "fmt"
    str "strings"
    msw "markw/lib"
    "math"
)

type Range struct {
    low int
    high int
}

type Mapping struct {
    src Range
    dest Range
}

type CategoryMap []Mapping

type Almanac struct {
    seeds []int
    maps  []CategoryMap
}

func (r Range) String() string {
    return fmt.Sprintf("%d-%d", r.low, r.high)
}

func (r Range) contains(n int) bool {
    return r.low <= n && n < r.high
}

func newMapping(ns []int) Mapping {
    dest := ns[0]
    src  := ns[1]
    length := ns[2]
    return Mapping{Range{src, src + length}, Range{dest, dest + length}}
}

func (m Mapping) mapSeed(seed int) int {
    if m.src.contains(seed) {
        newSeed := m.dest.low - m.src.low + seed
        return newSeed
    }
    return seed
}

func (sm CategoryMap) mapSeed(seed int) int {
    for _, r := range sm {
        val := r.mapSeed(seed)
        if (val != seed) {
            return val
        }
    }
    return seed
}

func (sm CategoryMap) mapSeedRange(r Range) Range {
    low := sm.mapSeed(r.low)
    high := low
    if r.low != r.high {
        high = sm.mapSeed(r.high - 1) + 1
    }
    return Range{low, high}
}

func parseAlmanac(ss []string) Almanac {
    var seeds []int
    var maps  []CategoryMap

    for _,s := range ss {
        if str.Index(s, "seeds:") == 0 {
            tokens := str.Fields(s)[1:]
            seeds = msw.Map(tokens, msw.ParseInt)
        } else if str.Index(s, "map:") > 1 {
            maps = append(maps, make([]Mapping, 0))
        } else if len(s) > 0 {
            ns := msw.Map(str.Fields(s), msw.ParseInt)
            index := len(maps)-1
            maps[index] = append(maps[index], newMapping(ns))
        }
    }
    return Almanac{seeds, maps}
}

func (a Almanac) seedRanges() []Range {
    result := []Range{}
    for i := 0; i < len(a.seeds); i += 2 {
        result = append(result, Range{a.seeds[i], a.seeds[i] + a.seeds[i+1]})
    }
    return result
}

func location(seed int, almanac Almanac) int {
    curr := seed
    for _, seedMap := range almanac.maps {
        next := seedMap.mapSeed(curr)
        curr = next
    }
    return curr
}

func minLocationPart1(almanac Almanac) int {
    minLoc := math.MaxInt32
    for _, seed := range almanac.seeds {
        loc := location(seed, almanac)
        if (loc < minLoc) {
            minLoc = loc
        }
    }
    return minLoc
}

func (r Range) partition(r0 Range) []Range {

    if (r.low >= r0.low && r.high <= r0.high) {
        return []Range{r}
    }

    if (r.low <= r0.low && r.high <= r0.high && r.high > r0.low) {
        return []Range{
            Range{r.low, r0.low}, 
            Range{r0.low, r.high}, 
        }
    }
    if (r.low >= r0.low && r.high >= r0.high && r.low < r0.high) {
        return []Range{
            Range{r.low,  r0.high}, 
            Range{r0.high, r.high}}
    }
    if (r.low <= r0.low && r.high >= r0.high) {
        return []Range{
            Range{r.low,   r0.low}, 
            Range{r0.low,  r0.high}, 
            Range{r0.high, r.high}}
    }
    return []Range{r}
}

func (cm CategoryMap) partitionRanges(rs []Range) []Range {
    result := msw.Reduce(cm, func (accum []Range, mapping Mapping) []Range {
        newRanges := []Range{}
        for _, r := range accum {
            partitions := r.partition(mapping.src)
            newRanges = append(newRanges, partitions...)
        }
        return newRanges
    }, rs)

    return msw.Map(result, cm.mapSeedRange)
}

func (cm CategoryMap) partition (r Range) []Range {
    return cm.partitionRanges([]Range{r})
}

func minLocationPart2(almanac Almanac) int {
    mappedSeedRanges := []Range{}
    for _, seedRange := range almanac.seedRanges() {

        mapped := msw.Reduce(almanac.maps, func (accum []Range, cm CategoryMap) []Range {
            return cm.partitionRanges(accum)
        }, []Range{seedRange})

        mappedSeedRanges = append(mappedSeedRanges, mapped...)
    }

    minLocation := math.MaxInt32
    for _, r := range mappedSeedRanges {
        if r.low < minLocation {
            minLocation = r.low
        }
    }
    return minLocation
}

func main() {
    lines := msw.FileLines("input.txt")
    almanac := parseAlmanac(lines)

    fmt.Printf("Part 1: %d\n", minLocationPart1(almanac))
    fmt.Printf("Part 2: %d\n", minLocationPart2(almanac))
}
