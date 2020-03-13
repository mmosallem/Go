package main

import (
	"fmt"
	"sort"
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
