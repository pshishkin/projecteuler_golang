package main

import (
	"log"
	"testing"
)

func TestSameAnswers(t *testing.T) {
	for N := 1; N <= 1000; N += 1 {
		var slowAnsList []ABC
		slowAns, slowErr := solveForParticularNSlow(N, &slowAnsList)
		if slowErr != nil {
			t.Fatalf("Error in SlowAns on N = %v: %v", N, slowErr)
		}
		var fastAnsList []ABC
		fastAns, fastErr := solveForParticularNFast(N, &fastAnsList)
		if fastErr != nil {
			t.Fatalf("Error in SlowAns on N = %v: %v", N, fastErr)
		}
		if slowAns != fastAns {
			log.Printf("Slow verbose answers:\n%v\n", slowAnsList)
			log.Printf("Fast verbose answers:\n%v\n", fastAnsList)
			t.Fatalf("Slow (%v) and Fast (%v) solvers return different answers on N = %v", slowAns, fastAns, N)
		}
	}
}
