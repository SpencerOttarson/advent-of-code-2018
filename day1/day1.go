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

    var values []int

    var firstDuplicate int = 0
    firstDuplicateFound := false
    counts := make(map[int]int)
    counts[0] = 1

    for scanner.Scan() {
        value, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        values = append(values, value)

        total += value

        value, ok := counts[total]
        if !ok {
            counts[total] = 1
        } else {
            if value >= 1 && firstDuplicateFound == false {
                firstDuplicate = total
                firstDuplicateFound = true
            }
            counts[total] = value + 1
        }
    }

    fmt.Println(total)

    for i := 0; firstDuplicateFound == false; i = (i + 1) % len(values) {
        total += values[i]

        value, ok := counts[total]
        if !ok {
            counts[total] = 1
        } else {
            if value >= 1 && firstDuplicateFound == false {
                firstDuplicate = total
                firstDuplicateFound = true
            }
            counts[total] = value + 1
        }
    } 
    
    fmt.Println(firstDuplicate)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
