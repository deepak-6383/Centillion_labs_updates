package main

import "fmt"

func main() {
    fmt.Print("Enter two numbers: ")
    var a, b int
    fmt.Scanln(&a, &b)
    fmt.Print("Enter operation (+, -, *, /): ")
    var op string
    fmt.Scanln(&op)
    switch op {
    case "+":
        fmt.Println("Result:", a+b)
    case "-":
        fmt.Println("Result:", a-b)
    case "*":
        fmt.Println("Result:", a*b)
    case "/":
        if b != 0 {
            fmt.Println("Result:", a/b)
        } else {
            fmt.Println("Cannot divide by zero")
        }
    default:
        fmt.Println("Invalid operator")
    }
}
