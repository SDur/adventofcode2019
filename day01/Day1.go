package main


import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    f, err := os.Open("input.txt")
    check(err)
    defer f.Close()

    var result int = 0

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        mass, _ := strconv.Atoi(scanner.Text())
        divided := mass / 3
        rounded := int(divided)
        substracted := rounded - 2
        result += substracted
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}