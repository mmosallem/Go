package main

import "fmt"

func main() {
	fmt.Println(121 % 10)
	fmt.Println("121123:", isPlaindromeNumber(567))
}
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
