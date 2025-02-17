package main

import "fmt"

func isPalindrome(s string) bool {
    rev := ""
    for i := len(s) - 1; i >= 0; i-- {
        rev += string(s[i])
    }
    return s == rev
}

func main() {
    fmt.Print("Enter a string: ")
    var input string
    fmt.Scanln(&input)
    if isPalindrome(input) {
        fmt.Println("It is a palindrome")
    } else {
        fmt.Println("Not a palindrome")
    }
}
