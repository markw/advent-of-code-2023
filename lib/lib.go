package lib 

import ( 
    "os" 
    str "strings"
    "reflect"
    "strconv"
)

func ParseInt(s string) int {
    n, err := strconv.Atoi(s)
    Check(err)
    return n
}

func Reduce[T, V any](ts []T, f func(V, T) V, initialValue V) V {
    accum := initialValue
    for _, t := range ts {
        accum = f(accum, t) 
    }
    return accum
}

func Check(err error) {
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

func FilterStr(s string, f func(int32) bool) string {
    return string(Filter([]rune(s), f))
}

func StrNotEmpty(s string) bool { 
    return len(s) > 0 
}

func FileLines(fileName string) []string {
    data, err := os.ReadFile(fileName)
    Check(err)
    return str.Split(string(data), "\n")
}

func SumInt(ns []int) int {
    addInt := func (a,b int) int { return a + b }
    return Reduce(ns, addInt, 0)
}

func Eq(a any, b any) bool {
    return reflect.DeepEqual(a,b)
}

type Integer interface {
    ~int | ~uint
}

func Min[T Integer] (ns []T) T {
    if len(ns) == 0 { panic("empty array") }
    min := ns[0]
    for _, n := range ns {
        if n < min {
            min = n
        }
    }
    return min
}

func Max[T Integer] (ns []T) T {
    if len(ns) == 0 { panic("empty array") }
    max := ns[0]
    for _, n := range ns {
        if n > max {
            max = n
        }
    }
    return max
}
