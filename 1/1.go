package main
import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Deret Fibonacci suku ke-0 sampai 10:")
	for i := 0; i < 11; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
}
