package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isNumPrime(23))
	fmt.Println(findAllPrimes(10))
}

func convertIntToFloat64(n int) float64 {
	return float64(n)
}
func isNumPrime(n int) bool {
	count := 1
	for i := 2; i <= int(math.Sqrt(convertIntToFloat64(n))); i++ {
		if n%i == 0 {
			count++
		}
	}
	return count != 2
}
func findAllPrimes(n int) []int {
	res := []int{1, 2}
	for p := 3; p <= n; p++ {
		count := 1
		for i := 2; i <= int(math.Sqrt(float64(p))); i++ {
			if p%i == 0 {
				fmt.Println("i:",i)
				count++
			}
		}
		if count == 2 {
			res = append(res, p)
		}
	}
	return res
}
