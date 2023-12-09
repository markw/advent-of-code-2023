package main

import (
    "fmt"
    "cmp"
    msw "markw/lib"
    str "strings"
    "strconv"
    "slices"
)

type Hand string

type Card rune

type CardMap map[rune]int

type CamelHand struct {
    hand Hand
    bid int
    score int
}

const (
    NO_PAIR = iota
    ONE_PAIR
    TWO_PAIR
    THREE_OF_A_KIND
    FULL_HOUSE
    FOUR_OF_A_KIND
    FIVE_OF_A_KIND
)
func (hand Hand) calcScore() int {
    cardMap := map[rune]int{}
    for _, c := range hand {
        cardMap[c] = cardMap[c] + 1
    }
    occurrences := make([]int, len(cardMap))
    i := 0
    for _, n := range cardMap {
        occurrences[i] = n
        i++
    }
    slices.Sort(occurrences)
    slices.Reverse(occurrences)
    //fmt.Printf("hand %s occ %s\n", hand, occurrences)
    if len(cardMap) == 5 { return NO_PAIR }        // 1 1 1 1 1
    if len(cardMap) == 4 { return ONE_PAIR }       // 2 1 1 1
    if len(cardMap) == 1 { return FIVE_OF_A_KIND } // 5
    if len(cardMap) == 3 {                         // 3 1 1 or 2 2 1
        if occurrences[0] == 3 { return THREE_OF_A_KIND }
        return TWO_PAIR
    }
    if len(cardMap) == 2 {  // 3 2 or 4 1
        if occurrences[0] == 3 { return FULL_HOUSE}
        return FOUR_OF_A_KIND
    }
    return len(cardMap)
}

func (c Card) cardValue() int {
    const cardChars = "23456789TJQKA"
    return str.IndexRune(cardChars, rune(c))
}

func compareCardValues(h0 Hand, h1 Hand) int {
    for i := 0; i < 5; i++ {
        c0 := Card(h0[i]).cardValue()
        c1 := Card(h1[i]).cardValue()
        if (c0 < c1) { return -1 }
        if (c0 > c1) { return  1 }
    }
    return 0
}

func parseCamelHand(s string) CamelHand {
    tokens := str.Fields(s)
    bid, err := strconv.Atoi(tokens[1])
    msw.Check(err)
    hand := Hand(tokens[0])
    return CamelHand{hand, bid, hand.calcScore()}
}

func main () {
    lines := msw.Filter(msw.FileLines("input.txt"), msw.StrNotEmpty)
    hands := msw.Map(lines, parseCamelHand)
    slices.SortFunc(hands, func (a,b CamelHand) int { 
        score := cmp.Compare(a.score, b.score)
        if score != 0 { return score }
        return compareCardValues(a.hand, b.hand)
    })

    winnings := 0
    for i, camelHand := range hands {
        // fmt.Println(camelHand)
        winnings += (i + 1) * camelHand.bid
    }
    fmt.Printf("Part 1: %d\n", winnings)
}
