package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var n, r int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	fmt.Print("Masukkan nilai r: ")
	fmt.Scan(&r)

	if n >= r {
		fmt.Printf("P(%d, %d) = %d\n", n, r, permutasi(n, r))
		fmt.Printf("C(%d, %d) = %d\n", n, r, combinasi(n, r))
	} else {
		fmt.Println("Nilai n harus lebih besar atau sama dengan r.")
	}
}
