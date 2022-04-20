package main

// https://projecteuler.net/problem=206

import (
	"errors"
	"math"
)

func isNumberFitsPattern(number int64) bool {
	var current_digit int64 = 0
	var current_base int64 = 1
	for pos := 0; pos <= 18; pos += 2 {
		number_digit_at_pos := number / current_base % 10
		if number_digit_at_pos != current_digit {
			return false
		}
		//update for the next cycle
		current_digit = (current_digit + 9) % 10
		current_base *= 100
	}
	return true
}

func Problem206() (int64, error) {
	const square_floor int64 = 1020304050607080900
	const square_ceil int64 = 1929394959697989990
	var ans_floor = int64(math.Sqrt(float64(square_floor)))
	var ans_ceil = int64(1. + math.Sqrt(float64(square_ceil)))

	for ans := ans_floor; ans < ans_ceil; ans += 1 {
		square := ans * ans
		if isNumberFitsPattern(square) {
			return ans, nil
		}
	}
	return 0, errors.New("not found an answer")
}
