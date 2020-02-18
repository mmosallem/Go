package main

import "fmt"

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
	fmt.Println("Majority item in ", arr, " result:", getMajorityItem(arr))

	fmt.Println("Fibonacci of 15:", getFibinacci(15))
	fmt.Println("IsSquere 2500:", isSquare(2500))

	var anotherArr = []int{10, 20, 30, 40, 50}
	fmt.Print("reversing array ", anotherArr)
	fmt.Print(" result:", reverseArray(anotherArr))
}