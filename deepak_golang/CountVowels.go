package main

import (
    "fmt"
    "strings"
)

func countVowels(s string) int {
    vowels := "aeiouAEIOU"
    count := 0
    for _, ch := range s {
        if strings.ContainsRune(vowels, ch) {
            count++
        }
    }
    return count
}

func main() {
    fmt.Print("Enter a string: ")
    var input string
    fmt.Scanln(&input)
    fmt.Println("Number of vowels:", countVowels(input))
}
