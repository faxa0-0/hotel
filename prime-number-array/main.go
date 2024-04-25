package main

import "fmt"

func main() {
	fmt.Println(findPrimes(11, 20))
}
func findPrimes(min, max int) []int {
	var primes []int
	
	for num := min; num <= max; num++ {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}

	return primes
}
func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
