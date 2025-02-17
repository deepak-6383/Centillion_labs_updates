// Simple Calculator
func calculator(a, b int, op string) int {
    switch op {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        if b != 0 {
            return a / b
        }
        fmt.Println("Cannot divide by zero")
    }
    return 0
}
