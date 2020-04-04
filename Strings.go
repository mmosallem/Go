package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	var r sortRunes = stringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

//reverseRunes reverse string
func reverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//isPermutation checkes if a is a permutaion of b
func isPermutation(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	a = sortStringByCharacter(a)
	b = sortStringByCharacter(b)
	return a == b
}

//printPermutaions prints all permutaions of a string
func printPermutaions(str string, prefix string) {
	if len(str) == 0 {
		fmt.Println(prefix)
	} else {
		for i := 0; i < len(str); i++ {
			restOfstring := string(str[0:i]) + string(str[i+1:])
			printPermutaions(restOfstring, prefix+string(str[i]))
		}
	}
}

//isBalanced checks if the string is balanced strings meaning [],{},() brackets are balanced
func isBalanced(str string) bool {
	var s StackString
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		var curr = string(r[i])
		if curr == "[" || curr == "{" || curr == "(" {
			s.Push(curr)
		} else if curr == "]" || curr == "}" || curr == ")" {
			var prev, b = s.Pop()
			if !b {
				return false
			}
			if (prev == "[" && curr != "]") || (prev == "{" && curr != "}") || (prev == "(" && curr != ")") {
				return false
			}

		}

	}
	if s.IsEmpty() {
		return true
	}
	return false
}

//printAllSubstrings prints all substrings of the supplied string
func printAllSubstrings(str string) {
	for l := 1; l < len(str); l++ {
		for i := 0; i <= len(str)-l; i++ {
			var s = str[i : i+l]
			fmt.Println(s)
		}
	}
}

func getAllSubstringsRemoving1Letter(str string) []string {
	var result []string

	for i := 0; i < len(str); i++ {
		result = append(result, string(str[0:i])+string(str[i+1:]))

	}
	return result
}

//compress string: aaaccddeff -> a3c2d2e1f2
func compress(str string) string {
	if len(str) == 0 {
		return ""
	}
	res := ""
	lastChar := str[0]
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] == lastChar {
			count++
		} else {
			res += string(lastChar) + strconv.Itoa(count)
			if len(res) > len(str) {
				return str
			}
			lastChar = str[i]
			count = 1
		}
	}
	res += string(lastChar) + strconv.Itoa(count)
	if len(res) > len(str) {
		return str
	}
	return res
}

//isPlaindrome checks whether the string is plaindrome or not
func isPlaindrome(str string) bool {
	return isPlaindromeString(str, 0, len(str)-1)
}

//isPlaindromeString checks whether the string is plaindrome or not
func isPlaindromeString(str string, start int, end int) bool {

	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}

//removeDupsChars removes dups characters from string
func removeDupsChars(str string) string {
	var result string = ""
	var table string = ""
	for i := 0; i < len(str); i++ {
		if strings.IndexByte(table, str[i]) == -1 {
			result += string(str[i])
			table += string(str[i])
		}

	}
	return result
}

//longestCommonSubstring returns the longest common substring for the supplied strings
func longestCommonSubstring(str1 string, str2 string) string {

	if len(str1) > len(str2) {
		return longestCommonSubstringInternal(str2, str1)
	}
	return longestCommonSubstringInternal(str1, str2)

}
func longestCommonSubstringInternal(small string, large string) string {
	if strings.Contains(large, small) {
		return small
	}
	if len(small) <= 1 {
		return ""
	}

	var result = ""
	var substr = getAllsubstringsOfLengthK(small, len(small)-1)

	for _, s := range substr {
		if strings.Contains(large, s) {
			if len(s) > len(result) {
				result = s
			}

		}
		var lcs = longestCommonSubstringInternal(s, large)
		if len(lcs) > len(result) {
			result = lcs
		}
	}

	return result
}

//getAllsubstringsOfLengthK gets all substrings of length k for the supplied string
func getAllsubstringsOfLengthK(str string, k int) []string {

	if str == "" {
		return []string{}
	}
	if k > len(str) {
		return []string{str}
	}

	size := len(str) - k + 1
	var res = make([]string, size)

	for i := 0; i < len(res); i++ {
		res[i] = str[i : k+i]
	}
	return res
}

//numberOfDeletionToGetCorrectWord returns the number of characters to be deleted from the supplied string to have a correct word that exists in the supplied dictionary
func numberOfDeletionToGetCorrectWord(word string, dictionary map[string]string) int {

	if dictionary[word] != "" {
		return 0
	}

	var queue []string
	queue = append(queue, word)

	for len(queue) > 0 {
		var str = queue[0]
		queue = queue[1:]

		if dictionary[str] != "" {
			return len(word) - len(str)
		}
		for i := 0; i < len(str); i++ {
			queue = append(queue, str[0:i]+str[i+1:])
		}
	}

	return -1

}
