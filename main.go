package main

import (
	"fmt"
	"hi-ch/hitachich"
)

func main() {
	minpairGot := hitachich.MinPair([]int{1, 4, 3, 2})
	palindromeGot := hitachich.Palindrome("aaa")
	numberCombineGot := hitachich.NumberCombine("23")

	fmt.Println("minpair input = [1, 4, 3, 2], got =", minpairGot)
	fmt.Println("palindrome input = aaa, got =", palindromeGot)
	fmt.Println("numberCombine input = 23, got =", numberCombineGot)
}
