package main
import "fmt"

func cetakganjil(n, x int) {
	if x > n {
		return
	}
	if x%2 != 0 {
		fmt.Print(x, " ")
	}
	cetakganjil(n, x+1)
}
func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan ganjil dari 1 hingga", n, "adalah:")
	cetakganjil(n, 1)
	fmt.Println()
}
