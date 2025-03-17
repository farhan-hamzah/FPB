// Pencari FPB
package main

import "fmt"

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var num1, num2 int

	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&num1)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&num2)

	fmt.Printf("FPB dari %d dan %d adalah: %d\n", num1, num2, gcd(num1, num2))
}

