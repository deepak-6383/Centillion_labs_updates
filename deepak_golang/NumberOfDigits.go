package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var num int
    fmt.Scanln(&num)
    count := 0
    temp := num
    for temp > 0 {
        temp /= 10
        count++
    }
    fmt.Println("Number of digits:", count)
}
