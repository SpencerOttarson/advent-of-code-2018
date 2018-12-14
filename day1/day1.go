package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("input1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    total := 0
    for scanner.Scan() {
        value, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        total += value
    }

    fmt.Println(total)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
