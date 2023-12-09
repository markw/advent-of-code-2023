package main

import (
    "fmt"
    msw "markw/lib"
)

type Steps struct {
    s string
    i int
}

func (steps *Steps) next() rune {
    if steps.i >= len(steps.s) { steps.i = 0 }
    next := rune(steps.s[steps.i])
    steps.i++
    return next
}

type Pair struct {
    left, right string
}

type Network map[string]Pair

func parseNetwork(ss []string) Network {
    network := Network{}
    for _, s := range ss {
        var name, left, right string
        fmt.Sscanf(s, "%s = (%3s, %3s)", &name, &left, &right)
        if len(name) > 0 {
            //fmt.Printf("name %s left %s right %s\n", name, left, right)
            network[name] = Pair{left, right}
        }
    }
    return network
}

func main () {
    lines := msw.FileLines("input.txt")
    steps := Steps{lines[0], 0}
    network := parseNetwork(lines[2:])

    curr := "AAA"
    count := 0
    for curr != "ZZZ" {
        step := steps.next()
        node := network[curr]
        next := node.left
        if step == 'R' { next = node.right }
        // fmt.Printf("node %s, step %s curr %s next %s count %d\n", node, string(step) ,curr, next, count)
        curr = next
        count++
    }
    fmt.Printf("Part 1: %d\n", count)
}
