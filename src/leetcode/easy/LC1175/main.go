package main

import "fmt"

func main() {
	fmt.Println(numPrimeArrangements(25) == 410206413)
	fmt.Println(numPrimeArrangements(15) == 261273600)
	fmt.Println(numPrimeArrangements(11) == 86400)
	fmt.Println(numPrimeArrangements(5) == 12)
	fmt.Println(numPrimeArrangements(100) == 682289015)
}
func numPrimeArrangements(n int) int {
	arr := make([]bool, n+1)
	count := (n + 1) >> 1
	NUM := 1000000007
	for i := 3; i*i <= n; i += 2 {
		for j := i * i; j <= n; j += 2 * i {
			if !arr[j] {
				arr[j] = true
				count--
			}
		}
	}
	prime := 1
	notPrime := 1
	for i := 2; i <= count; i++ {
		prime = (prime * i) % NUM
	}
	for i := 2; i <= n-count; i++ {
		notPrime = (notPrime * i) % NUM
	}
	return (prime * notPrime) % NUM
}
