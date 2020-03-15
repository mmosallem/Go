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

	fmt.Println("checking array is sorted ", anotherArr)
	fmt.Println("result:", isArraySortedMax1Swap(anotherArr))

	var str1 string = "abde"
	var str2 string = "bedav"
	fmt.Println("is permutation ", str1, " and ", str2, ": ", isPermutation(str1, str2))

	var strBalanceTEst = "ab(ccd)+[fdf]+{}}"
	fmt.Println("check if string is balanced ", strBalanceTEst, ":", isBalanced(strBalanceTEst))

	var testPerm string = "abc"
	fmt.Println("printing all permutations of ", testPerm)
	printPermutaions(testPerm, "")

	var testSubstring string = "abcde"
	fmt.Println("printing all subsctrings of ", testSubstring)
	printAllSubstrings(testSubstring)

	var testSubstringRemove1 = "abc"
	strs := getAllSubstringsRemoving1Letter(testSubstringRemove1)
	fmt.Println("printing all subsctrings (removing 1 letter) of ", testSubstringRemove1, ":", strs)

	var testCompress = "aaaccddeff"
	fmt.Println("compressing string ", testCompress, ". result:", compress(testCompress))

	var testPlaindromeStr = "abcba"
	fmt.Println("Is Plaindrome string ", testPlaindromeStr, ": ", isPlaindrome(testPlaindromeStr))
}
