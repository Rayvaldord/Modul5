package main
import "fmt"

func urutan(n, x int) {
	fmt.Print(x, " ")

	if x > 1 {
		urutan(n, x-1)
	}
	if x < n {
		fmt.Print(x+1, " ")
	}
}
func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Println("Hasil Urutan:")
	urutan(n, n)
	fmt.Println()
}
