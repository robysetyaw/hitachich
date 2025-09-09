package main

import "hi-ch/hitachich"

func main() {
	minpairGot := hitachich.MinPair([]int{1, 4, 3, 2})
	palindromeGot := hitachich.Palindrome("aaa")
	numberCombineGot := hitachich.NumberCombine("23")

	println("minpair input = [1, 4, 3, 2] got = ", minpairGot)
	println("palindrome input = aaa got = ", palindromeGot)
	println("numberCombine input = 23 got = ", numberCombineGot)
}
