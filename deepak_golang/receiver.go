// Receiver Functions with Pointers
package main
import "fmt"
type Counter struct {
    count int
}
func (c *Counter) Increment() {
    c.count++
}
func main() {
    c := Counter{count: 0}
    c.Increment()
    fmt.Println(c.count)
}