package main

import (
    "fmt"
)

type Race struct {
    time int
    distance int
}

func race(r Race) int {
    count := 0
    for i := 0; i < r.time; i++ {
        travelTime := r.time - i;
        distance := travelTime * i
        if distance > r.distance {
            count++
        }
    }
    return count;
}

func main() {
    // races := []Race{ Race{7,9}, Race{15,40}, Race{30, 200}}
    races := []Race{ Race{44,283}, Race{70,1134}, Race{70,1134}, Race{80,1491}}
    result := 1
    for _, r := range races {
        result *= race(r)
    }
    fmt.Printf("Part 1: %d\n", result)
    fmt.Printf("Part 2: %d\n", race(Race{44707080, 283113411341491}))
}
