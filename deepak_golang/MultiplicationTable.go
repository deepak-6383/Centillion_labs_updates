package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var num int
    fmt.Scanln(&num)
    fmt.Println("Multiplication table for", num)
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", num, i, num*i)
    }
}
