package main

// https://projecteuler.net/problem=474

import (
	"math"
	"sort"
)

const (
	M int64 = 10*1000*1000*1000*1000*1000 + 61
)

func getPrimeFactorisation(N int) []int {
	ans := make([]int, 0)
	sqr := int(math.Sqrt(float64(N) + 1))
	for i := 2; i <= sqr; i += 1 {
		for N%i == 0 {
			N /= i
			ans = append(ans, i)
		}
	}
	if N > 1 {
		ans = append(ans, N)
	}
	return ans
}

func getLastDigitsMod(lastDigits int) int {
	ans := 1
	for lastDigits > 0 {
		ans *= 10
		lastDigits /= 10
	}
	return ans
}

func getBruteforceAns(nFactorial int, lastDigits int) int64 {
	lastDigitsMod := getLastDigitsMod(lastDigits)

	var factorial int = 1
	for i := 2; i <= nFactorial; i += 1 {
		factorial *= i
	}

	var ans int64
	for i := 1; i <= factorial; i += 1 {
		if factorial%i == 0 && i%lastDigitsMod == lastDigits {
			ans = (ans + 1) % M
		}
	}

	return ans
}

func getFastAns(nFactorial int, lastDigits int) int64 {
	allDivisors := make([]int, 1)
	allDivisors[0] = 1

	for i := 2; i <= nFactorial; i += 1 {
		divisors := getPrimeFactorisation(i)
		allDivisors = append(allDivisors, divisors...)
	}

	sort.Ints(allDivisors)

	lastDigitsMod := getLastDigitsMod(lastDigits)

	ans := make([]int64, lastDigitsMod)

	ans[1] = 1

	for divisorIntervalStart := 1; divisorIntervalStart < len(allDivisors); {
		divisorIntervalEnd := divisorIntervalStart + 1
		for divisorIntervalEnd < len(allDivisors) && allDivisors[divisorIntervalEnd] == allDivisors[divisorIntervalStart] {
			divisorIntervalEnd += 1
		}
		currentDivisor := allDivisors[divisorIntervalStart]
		sameDivisors := divisorIntervalEnd - divisorIntervalStart

		ansNext := make([]int64, len(ans))

		var divMul int64 = 1
		for divPower := 0; divPower <= sameDivisors; divPower += 1 {
			for t := 0; t < lastDigitsMod; t += 1 {
				ansNext[divMul*int64(t)%int64(lastDigitsMod)] += ans[t]
				ansNext[divMul*int64(t)%int64(lastDigitsMod)] %= M
			}
			divMul = divMul * int64(currentDivisor) % int64(lastDigitsMod)
		}

		ans = ansNext
		divisorIntervalStart = divisorIntervalEnd
	}
	return ans[lastDigits]
}

func Problem474() (int, error) {
	//return int(getFastAns(1000000, 65)), nil
	return int(getFastAns(1000000, 65432)), nil
}
