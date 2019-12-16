package main


import (
    "fmt"
    "strconv"
)


func main() {
    start, end := 156218, 652527
    possibilities := 0

    for i := start; i <= end; i++ {
        increasing := isIncreasing(i)
        double := hasDouble(i)
        if increasing && double {
            possibilities++
        }
    }
    fmt.Printf("And the answer: %d", possibilities)
}

func isIncreasing(input int) bool {
    stringInput := strconv.Itoa(input)
    for i := 1; i < len(stringInput); i++ {
        if int(stringInput[i -1]) > int(stringInput[i]) {
            return false
        }
    }
    return true
}

func hasDouble(input int) bool {
    stringInput := strconv.Itoa(input)
    for i := 1; i < len(stringInput); i++ {
        if int(stringInput[i -1]) == int(stringInput[i]) {
            return true
        }
    }
    return false
}