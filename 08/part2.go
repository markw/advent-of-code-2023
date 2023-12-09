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

func parseNetwork(ss []string) (Network, []string, []string) {
    network := Network{}
    startingNodes := []string{}
    endingNodes := []string{}
    for _, s := range ss {
        var name, left, right string
        fmt.Sscanf(s, "%s = (%3s, %3s)", &name, &left, &right)
        if len(name) > 0 {
            //fmt.Printf("name %s left %s right %s\n", name, left, right)
            network[name] = Pair{left, right}
            if name[2] == 'A' {
                startingNodes = append(startingNodes, name)
            }
            if name[2] == 'Z' {
                endingNodes = append(endingNodes, name)
            }
        }
    }
    return network, startingNodes, endingNodes
}

func gcd(a,b int) int {
    if b == 0 { return a }
    return gcd(b, a % b)
}

func lcm(ns []int) int {
    return msw.Reduce(ns, func(a,b int) int { return a * b / gcd(a,b) }, 1)
}

func traverse(start, end string, network Network, steps Steps) int {
    result := 1
    curr := start
    visited := make(map[string]bool)
    for curr != end {
        step := steps.next()
        node := network[curr]
        next := node.left
        if step == 'R' { next = node.right }
        if next == end {
            return result
        }
        // cycle detection
        key := fmt.Sprintf("%s-%d", next, result % len(steps.s))
        if visited[key] == true {
            // fmt.Printf("%s -> %s: cycle at step %d, key %s\n", start, end, result, key)
            break
        }
        visited[key] = true
        // end cycle detection

        curr = next
        result++
    }
    return 0
}

func main () {
    lines := msw.FileLines("input.txt")
    steps := Steps{lines[0], 0}
    // fmt.Printf("Steps length: %d\n", len(steps.s))
    network, startingNodes, endingNodes := parseNetwork(lines[2:])
    fmt.Printf("starting:     %s\n", startingNodes);
    fmt.Printf("ending:       %s\n", endingNodes);

    distances := []int{}
    for _, startNode := range startingNodes {
        for _, endNode := range endingNodes {
            distance := traverse(startNode,endNode, network, steps)
            if (distance > 0) {
                fmt.Printf("%s -> %s: %d\n", startNode, endNode, distance)
                distances = append(distances, distance)
            }
        }
    }

    // fmt.Printf("Distances:    %v\n", distances)
    // fmt.Printf("Modulos:      %v\n", msw.Map(distances,func (n int) int { return n % len(steps.s) }))
    fmt.Printf("Part 2:       %d\n", lcm(distances))
}
