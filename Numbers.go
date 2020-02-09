package main

import "fmt"

func main() {

	fmt.Println("121123:", isPlaindromeNumber(567))

	fmt.Println("factorial for 12: ", factorial(12))
}

// isPlaindromeNumber returns whether a number is plaindrom number or not
func isPlaindromeNumber(x int64) bool {
	return reverseNumber(x) == x
}
func reverseNumber(number int64) int64 {
	var retVal int64 = 0
	for {
		retVal = (retVal * 10) + (number % 10)
		number = number / 10
		if number == 0 {
			break
		}
	}
	return retVal
}

//  factorial returns the factorial for the given number
func factorial(x int64) int64 {
	if x <= 1 {
		return 1
	}
	var n int64
	n = 1

	var i int64
	for i = 2; i <= x; i++ {
		n = i * n
	}
	return n
}
