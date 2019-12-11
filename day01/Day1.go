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
        for {
            divided := mass / 3
            rounded := int(divided)
            substracted := rounded - 2
            if (substracted <= 0) {
                break
            }
            result += substracted
            mass = substracted
        }
    }

    fmt.Println(result)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}