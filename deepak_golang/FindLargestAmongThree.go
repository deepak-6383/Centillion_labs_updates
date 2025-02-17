package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Print("Enter three numbers: ")
    fmt.Scanln(&a, &b, &c)
    if a >= b && a >= c {
        fmt.Println("Largest number:", a)
    } else if b >= a && b >= c {
        fmt.Println("Largest number:", b)
    } else {
        fmt.Println("Largest number:", c)
    }
}
