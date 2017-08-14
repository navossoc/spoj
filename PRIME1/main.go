package main

import "fmt"
import "math"

func main() {
	var t int
	var m, n int

	fmt.Scanln(&t)
	for ; t > 0; t-- {
		fmt.Scanln(&m, &n)
		for ; m <= n; m++ {
			if isPrime(m) {
				fmt.Println(m)
			}
		}
		fmt.Println()
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n != 2 && n%2 == 0 {
		return false
	}

	// Trial division
	r := int(math.Sqrt(float64(n)))

	for i := 3; i <= r; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
