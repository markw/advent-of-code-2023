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

func processCards(originals, cards Cards) Cards {
    var copies Cards
    for _, card := range cards {
        if card.matches > 0 {
            i := card.id
            j := i + card.matches
            copies = append(copies, originals[i:j]...)
        }
    }
    return copies
}

func main() {
    lines := msw.Filter(msw.FileLines("input.txt"), msw.StrNotEmpty)
    cards := msw.Map(lines, parseCard)
    total := msw.SumInt(msw.Map(cards, score))
    fmt.Printf("Part 1: %d\n", total)

    copies := cards
    count := len(copies)
    for len(copies) > 0 {
        copies = processCards(cards, copies)
        count += len(copies)
    }
    fmt.Printf("Part 2: %d\n", count)
}


