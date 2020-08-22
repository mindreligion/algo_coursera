package prime

import "fmt"

func PrimeN(n int) int {
	if n <= 1 {
		return 1
	}
	var primes []int
outer:
	for i := 2; i < n; i++ {
		for _, j := range primes {
			if i%j == 0 {
				continue outer
			}
		}

		primes = append(primes, i)
		fmt.Println("prime processed", i)
	}

	if len(primes) == 0 {
		return 1
	}

	return primes[len(primes)-1]
}
