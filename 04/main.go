package main

import (
    msw "markw/lib"
    "fmt"
    str "strings"
    "strconv"
    "slices"
)

type Card struct {
    id int
    winners []string
    numbers []string
    matches int
}

type Cards []Card

func isNumber(s string) bool {
    _, err := strconv.Atoi(s)
    if (err != nil) { return false }
    return true
}

func parseCard(s string) Card {
    winners := []string{}
    numbers := []string{}
    mode := 0
    var cardNum int
    for _, token := range str.Fields(s) {
        if isNumber(token) {
            if mode == 0 {
                winners = append(winners, token)
            } else {
                numbers = append(numbers, token)
            }
        } else if token == "|" {
            mode = 1
        } else if str.Contains(token, ":") {
            cardNum, _ = strconv.Atoi(str.Trim(token, ":"))
        }
    }
    matches := 0
    for _, w := range winners {
        if slices.Contains(numbers, w) {
            matches++
        }
    }
    return  Card{cardNum, winners, numbers, matches}
}

func score(card Card) int {
    if (card.matches == 0) { return 0 }
    score := 1
    for i := 0; i < card.matches - 1; i++ {
        score *= 2
    }
    return score
}

type Node struct {
    value int
    next *Node
}

func addNode(tail *Node, n int) *Node {
    next := &Node{n, nil}
    tail.next = next
    return next
}

func processCards(cards Cards) int {
    head := &Node{-1, nil}
    tail := head
    for _, card := range cards {
        tail = addNode(tail, card.id - 1)
    }
    count := len(cards)
    curr := head.next
    for curr != nil {
        card := cards[curr.value]
        if card.matches > 0 {
            i := card.id
            j := i + card.matches
            for k := i; k < j; k++ {
                tail = addNode(tail, k)
                count++
            }
        }
        curr = curr.next
    }
    return count
}

func main() {
    lines := msw.Filter(msw.FileLines("input.txt"), msw.StrNotEmpty)
    cards := msw.Map(lines, parseCard)
    total := msw.SumInt(msw.Map(cards, score))
    fmt.Printf("Part 1: %d\n", total)
    fmt.Printf("Part 2: %d\n", processCards(cards))
}
