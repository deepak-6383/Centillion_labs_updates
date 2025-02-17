package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var num int
    fmt.Scanln(&num)
    isPrime := true
    if num < 2 {
        isPrime = false
    } else {
        for i := 2; i*i <= num; i++ {
            if num%i == 0 {
                isPrime = false
                break
            }
        }
    }
    if isPrime {
        fmt.Println(num, "is a Prime number")
    } else {
        fmt.Println(num, "is not a Prime number")
    }
}
