package main
import "fmt"
func main() {
    arr := []int{10, 20, 30, 40, 50}
    fmt.Println("Original array:", arr)
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    fmt.Println("Reversed array:", arr)
}
