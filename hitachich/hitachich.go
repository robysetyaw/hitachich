package hitachich

import (
	"sort"
)

var numberCharsMap = map[string]string{
	"0": "",
	"1": "",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func MinPair(input []int) int {

	sort.Ints(input)

	sum := 0
	for i := 0; i < len(input); i += 2 {
		sum += input[i]
	}

	return sum
}

func Palindrome(input string) int {
	n := len(input)
	count := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			substring := input[i : j+1]
			if isPalindrome(substring) {
				count++
			}
		}
	}

	return count
}

func isPalindrome(substring string) bool {
	left := 0
	right := len(substring) - 1
	for left < right {
		if substring[left] != substring[right] {
			return false
		}
		right--
		left++
	}
	return true
}

func NumberCombine(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	result := []string{""}

	for i := 0; i < len(input); i++ {
		number := string(input[i])
		char := numberCharsMap[number]
		var temp []string

		for _, prefix := range result {
			for j := 0; j < len(char); j++ {
				// fmt.Println(temp, "-", prefix, "-", letters[j])
				temp = append(temp, prefix+string(char[j]))
			}
		}
		result = temp
	}

	return result
}
