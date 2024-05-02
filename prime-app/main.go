package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	fmt.Println(isPrime(3))
	fmt.Println(isPrime(30))

}

func isPrime(number int) (bool, string) {
	// isPrime := true
	if number == 0 || number == 1 {
		return false, fmt.Sprintf("%d is not primenumber by definition", number)
	}
	if number < 0 {
		return false, "negative numbers are not prime number by definesiton"
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number as it is divisible by %d", number, i)
		}
	}
	return true, fmt.Sprintf("Bingo!! %d is a prime number", number)
}
