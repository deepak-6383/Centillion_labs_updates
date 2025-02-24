package main
import (
    "fmt"
    "sort"
)

func main() {
    var n int
    fmt.Print("Enter the size of the array: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    fmt.Println("Enter", n, "numbers:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    sort.Ints(arr)

    fmt.Println("Sorted array:", arr)
}
