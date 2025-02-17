package main

import "fmt"

func fibonacci(n int) []int {
    seq := []int{0, 1}
    for i := 2; i < n; i++ {
        seq = append(seq, seq[i-1]+seq[i-2])
    }
    return seq
}

func main() {
    fmt.Print("Enter number of Fibonacci terms: ")
    var n int
    fmt.Scanln(&n)
    fmt.Println("Fibonacci sequence:", fibonacci(n))
}
