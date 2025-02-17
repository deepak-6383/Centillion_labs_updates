// Parsing Floats
package main
import (
    "fmt"
    "strconv"
)
func main() {
    num, _ := strconv.ParseFloat("3.14", 64)
    fmt.Println(num)
}