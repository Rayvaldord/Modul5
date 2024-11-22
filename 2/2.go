package main
import "fmt"

func bintang(n int) {
	if n <= 0 {
		return
	}
	bintang(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	fmt.Println("Pola bintang:")
	bintang(n)
}
