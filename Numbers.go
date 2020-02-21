package main

import (
	"errors"
	"math"
)

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
	}
	return false

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

//getFibinacci
func getFibinacci(x int) int {
	var arr = make([]int, x)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < x; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[x-1] + arr[x-2]
}

//isSquare returns whether the number is square number of not
func isSquare(n int) bool {
	if n < 0 {
		n = -n
	}

	if n <= 1 {
		return true
	}
	var currSq = 4
	var currN = 2

	for currSq <= n {
		if currSq == n {
			return true
		}
		currN++
		currSq = currN * currN

	}
	return false
}
