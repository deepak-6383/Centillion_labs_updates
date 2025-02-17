package main

import "fmt"

func main() {
    fmt.Print("Enter two numbers: ")
    var a, b int
    fmt.Scanln(&a, &b)
    fmt.Println("Sum:", a+b)
}
