package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Print a welcome message
	intro()

	// Create a channel to indicate you want to quit
	doneChan := make(chan bool)

	// start go routine to read user nput
	go readUserInput(doneChan)

	// block until channel gets a value
	<-doneChan
	// close the channel
	close(doneChan)
	fmt.Println("bye")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumber(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		promt()
	}

}

func checkNumber(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	} else {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return "NaN, please input a number", false
		} else {
			_, result := isPrime(num)
			return result, false
		}

	}
}

func intro() {
	fmt.Println("Hello...")
	fmt.Println("Enter any number to see if it's prime or not! Press q to quit.")
	promt()
}
func promt() {
	fmt.Print("-> ")
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
