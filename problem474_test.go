package main

import (
	"reflect"
	"testing"
)

func TestPrimeFactorisation(t *testing.T) {

	if reflect.DeepEqual(getPrimeFactorisation(35), []int{5, 7}) == false {
		t.Fatalf("Wrong factorisation of = %v: %v", getPrimeFactorisation(35), []int{5, 7})
	}
	if reflect.DeepEqual(getPrimeFactorisation(37), []int{37}) == false {
		t.Fatalf("Wrong factorisation of = %v: %v", getPrimeFactorisation(35), []int{37})
	}

}

func TestBruteForceSolution(t *testing.T) {

	if getBruteforceAns(12, 12) != 11 {
		t.Fatalf("F(12!, 12) = %v, but correct answer is %v", getBruteforceAns(12, 12), 11)
	}

}

func TestFastSolutionWithBruteForce(t *testing.T) {
	for nFact := 1; nFact <= 12; nFact += 1 {
		for divLastDigits := 1; divLastDigits <= 11; divLastDigits += 1 {
			bruteAns := getBruteforceAns(nFact, divLastDigits)
			fastAns := getFastAns(nFact, divLastDigits)
			if bruteAns != fastAns {
				t.Fatalf("F(%v!, %v) = %v, but correct answer is %v", nFact, divLastDigits, bruteAns, fastAns)
			}
		}
	}
}

func BigNumberTest(t *testing.T) {
	nFact := 100000
	divLastDigits := 65
	fastAns := getFastAns(nFact, divLastDigits)
	rightAns := 7993346047434016
	if fastAns != 7993346047434016 {
		t.Fatalf("F(%v!, %v) = %v, but correct answer is %v", nFact, divLastDigits, fastAns, rightAns)
	}

}
