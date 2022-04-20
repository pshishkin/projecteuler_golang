package main

import (
	"errors"
	"math"
	"runtime/debug"
)

// https://projecteuler.net/problem=206

func generateIsCompositeArray(max_n int) []byte {
	isComposite := make([]byte, max_n)
	for i := 2; i*i < max_n; i += 1 {
		if isComposite[i] == 1 {
			continue
		}
		for j := 2; i*j < max_n; j += 1 {
			isComposite[i*j] = 1
		}
	}
	return isComposite
}

func getPrimeNumbersList(max_n int) []int {
	isComposite := generateIsCompositeArray(max_n)
	var primeNumberList []int
	for key, value := range isComposite {
		if value == 0 && key >= 2 {
			primeNumberList = append(primeNumberList, key)
		}
	}
	return primeNumberList
}

type ABC struct {
	a int64
	b int64
	c int64
}

func validateAB(a int64, b int64, sqrt int64, N int) int64 {
	if a <= b && b <= int64(N) {
		numerator := a * b
		denominator := 2*sqrt + a + b
		if numerator%denominator == 0 {
			return numerator / denominator
		}
	}
	return -1
}

// move some variables outside to save stack size
var ()

func iterateAfterAB(a int64, b int64, sqrt int64, pIndex int, N int, primeNumberList *[]int, ansList *[]ABC) (int64, error) {
	if pIndex == len(*primeNumberList) {
		return 0, nil
	}

	currentPrime := (*primeNumberList)[pIndex]

	var localAns int64 = 0

	// Check that further iterations are needed at all, paying attention to overflow, i.e.
	// N = 10^9
	// a and b ~= 10^9
	if a*int64(currentPrime) > int64(N) || b*int64(currentPrime) > int64(N) {
		return localAns, nil
	}

	primePowerA := 0
	primeMulA := int64(1)
	sqrtMulA := int64(1)
	for a*primeMulA <= int64(N) {
		primePowerB := 0
		primeMulB := int64(1)
		sqrtMulB := int64(1)

		if primePowerA%2 == 1 {
			primePowerB += 1
			primeMulB *= int64(currentPrime)
			sqrtMulB *= int64(currentPrime)
		}
		for b*primeMulB <= int64(N) {

			if primePowerA != 0 || primePowerB != 0 {
				// something new and unique, let's validate it
				c := validateAB(a*primeMulA, b*primeMulB, sqrt*sqrtMulA*sqrtMulB, N)
				if c != -1 {
					localAns += a*primeMulA + b*primeMulB + c
					*ansList = append(*ansList, ABC{a, b, c})
				}
			}

			retAns, err := iterateAfterAB(a*primeMulA, b*primeMulB, sqrt*sqrtMulA*sqrtMulB, pIndex+1, N, primeNumberList, ansList)
			if err != nil {
				return retAns + localAns, err
			}
			localAns += retAns

			// Prevent overflow
			if int64(currentPrime)*int64(currentPrime) > int64(N) {
				break
			}
			primePowerB += 2
			primeMulB *= int64(currentPrime) * int64(currentPrime)
			sqrtMulB *= int64(currentPrime)
		}
		// prepare for the next iteration
		primePowerA += 1
		primeMulA *= int64(currentPrime)
		if primePowerA%2 == 0 {
			sqrtMulA *= int64(currentPrime)
		}
	}
	return localAns, nil
}

func solveForParticularNFast(N int, ansList *[]ABC) (int64, error) {
	if N > (1 << 20) {
		debug.SetMaxStack(1 << 33)
	}

	primeNumberList := getPrimeNumbersList(N)
	// fmt.Println(N, len(primeNumberList))
	// fmt.Println(N/len(primeNumberList), math.Log(float64(N)))

	return iterateAfterAB(1, 1, 1, 0, N, &primeNumberList, ansList)
}

func solveForParticularNSlow(N int, ansList *[]ABC) (int64, error) {
	if N < 1 || N > 1000 {
		return 0, errors.New("slow solver only supports N from 1 to 1000")
	}
	var localAns int64
	for b := 1; b <= N; b += 1 {
		for a := 1; a <= b; a += 1 {
			numerator := a * b
			abSqrt := int(math.Sqrt(float64(a*b) + 0.1))
			if abSqrt*abSqrt != numerator {
				continue
			}
			denominator := 2*abSqrt + a + b
			if numerator%denominator == 0 {
				c := numerator / denominator
				*ansList = append(*ansList, ABC{int64(a), int64(b), int64(c)})
				localAns += int64(a + b + c)
			}
		}
	}
	return localAns, nil
}

func Problem510() (int64, error) {
	// ans := solveForParticularN(1000)
	// fmt.Println(ans)
	ansList := make([]ABC, 0)
	//return solveForParticularNFast(3, &ansList)
	return solveForParticularNFast(100000000, &ansList)
	// return solveForParticularNFast(100, &ansList)
	//return 0, errors.New("not found an answer")
}
