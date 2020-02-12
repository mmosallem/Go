package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	fmt.Println("121123:", isPlaindromeNumber(567))

	x, err := factorial(12)
	if err != nil {
		fmt.Println("error calculating factorial 12")
	} else {
		fmt.Println("factorial for 12: ", x)
	}

	fmt.Println("is this a prime number 13:", isPrime(13))

	fmt.Println("Is Armstrong 153:", isArmstrongNumber(153))

	var arr = []int{10, 20, 30, 20, 10, 10, 10}
	fmt.Println("Majority item in {10, 20, 30, 20, 10,10,10}:", getMajorityItem(arr))
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
func factorial(x int64) (int64, error) {
	if x < 0 {
		return -1, errors.New("cannot calculate factorial for negative numbers")
	}
	if x <= 1 {
		return 1, nil
	}
	var n int64
	n = 1

	var i int64
	for i = 2; i <= x; i++ {
		n = i * n
	}
	return n, nil
}

// isPrime checks whether the supplier number is integer or not
func isPrime(x int) bool {
	var i int
	for i = 2; i < int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

//isArmstrongNumber checks if a number is armstrong number
//An Armstrong number is a number such that the sum of its digits raised to the third power is equal to the number itself.
func isArmstrongNumber(x int) bool {
	var sum int = 0
	var remainder int = 0
	for i := x; i > 0; i = i / 10 {
		remainder = i % 10
		sum += remainder * remainder * remainder
	}
	if sum == x {
		return true
	} else {
		return false
	}
}

//getMajorityItem returns the majority item in array
func getMajorityItem(arr []int) int {

	var m = make(map[int]int)
	var majority int = len(arr) / 2
	for _, x := range arr {
		_, ok := m[x]
		if ok {
			m[x]++
			if m[x] > majority {
				return x
			}
		} else {
			m[x] = 1
		}
	}

	return -1
}
