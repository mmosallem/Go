package main

import "sort"

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
