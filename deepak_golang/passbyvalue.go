// Pass by Value
package main
import "fmt"
func modify(x int) {
    x = 20
}
func main() {
    num := 10
    modify(num)
    fmt.Println(num)
}
