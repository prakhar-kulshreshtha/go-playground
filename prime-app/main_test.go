package main

import "testing"

func Test_isPrime(t *testing.T) {
	prime_tests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "Bingo!! 7 is a prime number"},
		{"prime", 8, false, "8 is not a prime number as it is divisible by 2"},
	}
	// result, msg := isPrime(0)
	// if result {
	// 	t.Errorf("0 is not a prime number")
	// }
	// if msg != "0 is not primenumber by definition" {
	// 	t.Errorf("Message didn't match")
	// }
	for _, e := range prime_tests {
		result, msg := isPrime(e.testNum)
		if e.expected != result {
			t.Errorf("expected %v, got %v", e.expected, result)
		}
		if e.msg != msg {
			t.Errorf("expected: %s , got %s", e.msg, msg)
		}

	}
}
