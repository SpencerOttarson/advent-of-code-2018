package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func differences(string1, string2 string) int {
    if len(string1) != len(string2) {
        return len(string1)
    }

    differences := 0
    for i, char := range string1 {
        if string(string2[i]) != string(char) {
            differences++
        }
    }

    return differences
}

func commonChars(string1, string2 string) string {
    if len(string1) != len(string2) {
        return ""
    }

    for i, char := range string1 {
        if string(string2[i]) != string(char) {
            return string1[0:i] + string1[i+1:len(string1)]
        }
    }

    return ""
}

func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    doubleCount := 0
    tripleCount := 0

    var words []string

    for scanner.Scan() {
        word := scanner.Text()
        words = append(words, word)

        counts := make(map[rune]int)
        for _, char := range word {
            count, ok := counts[char]
            if !ok {
                counts[char] = 1
            } else {
                counts[char] = count + 1
            }
        }

        hasDouble := 0
        hasTriple := 0
        for _, value := range counts {
            if value == 2 {
                hasDouble = 1
            } else if value == 3 {
                hasTriple = 1
            }
        }
        doubleCount += hasDouble
        tripleCount += hasTriple
    }

    fmt.Println(doubleCount * tripleCount)

    for i, word := range words {
        for j := i + 1; j < len(words); j++ {
            if differences(word, words[j]) == 1 {
                fmt.Println(commonChars(word, words[j]))
            }
        }
    }
}
