package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var num int
    fmt.Scanln(&num)
    if num%2 == 0 {
        fmt.Println(num, "is Even")
    } else {
        fmt.Println(num, "is Odd")
    }
}
