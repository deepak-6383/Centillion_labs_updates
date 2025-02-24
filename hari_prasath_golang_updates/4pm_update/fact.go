package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	result := factorial(n)
	fmt.Println(result)
}
