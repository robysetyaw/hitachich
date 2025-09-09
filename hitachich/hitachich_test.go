package hitachich

import (
	"reflect"
	"testing"
)

func TestHitachi(t *testing.T) {
	runTestingMinPair(t)
	runTestingNumberCombine(t)
	runTestingPalindrome(t)
}

func runTestingMinPair(t *testing.T) {
	runTestMinPair(t, "should return 1 when input is [1, 2]", []int{1, 2}, 1)
	runTestMinPair(t, "should return 4 when input is [1, 4, 3, 2]", []int{1, 4, 3, 2}, 4)
}

func runTestingNumberCombine(t *testing.T) {
	runTestNumberCombine(t, "should return [] when input is 0", "0", nil)
	runTestNumberCombine(t, "should return [] when input is 10", "10", nil)
	runTestNumberCombine(t, "should return ['a', 'b', 'c'] when input is 2", "2", []string{"a", "b", "c"})
	runTestNumberCombine(t, "should return ['a', 'b', 'c'] when input is 20", "20", nil)
	runTestNumberCombine(t, "should return ['ad', 'ae', 'af', 'bd', 'be', 'bf', 'cd', 'ce', 'cf'] when input is 23", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
}

func runTestingPalindrome(t *testing.T) {
	runTestPalindrome(t, "should return 1 when input is 'a'", "a", 1)
	runTestPalindrome(t, "should return 2 when input is 'ab'", "ab", 2)
	runTestPalindrome(t, "should return 6 when input is 'aaa'", "aaa", 6)
	runTestPalindrome(t, "should return 7 when input is 'katak'", "katak", 7)
}

func runTestMinPair(t *testing.T, description string, input []int, expected int) {
	t.Run(description, func(t *testing.T) {
		got := MinPair(input)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Failed : %s, minPair(%v) = %v; want %v", description, input, got, expected)
		}
	})
}

func runTestNumberCombine(t *testing.T, description string, input string, expected []string) {
	t.Run(description, func(t *testing.T) {
		got := NumberCombine(input)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Failed : %s, numberCombine(%v) = %v; want %v", description, input, got, expected)
		}
	})
}

func runTestPalindrome(t *testing.T, description string, input string, expected int) {
	t.Run(description, func(t *testing.T) {
		got := Palindrome(input)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Failed : %s, isPalindrome(%v) = %v; want %v", description, input, got, expected)
		}
	})

}
