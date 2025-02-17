package main

import "fmt"

func reverse(n int) int {
    rev := 0
    for n > 0 {
        rev = rev*10 + n%10
        n /= 10
    }
    return rev
}

func main() {
    fmt.Print("Enter a number: ")
    var num int
    fmt.Scanln(&num)
    fmt.Println("Reversed number:", reverse(num))
}
