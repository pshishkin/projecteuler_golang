package main

import (
	"log"
	"sort"
	"testing"
)

func TestBruteforceProblem510(t *testing.T) {
	for N := 1; N <= 3000; N += 1 {
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
			sort.Slice(fastAnsList, func(i, j int) bool {
				if fastAnsList[i].b != fastAnsList[j].b {
					return fastAnsList[i].b < fastAnsList[j].b
				} else {
					return fastAnsList[i].a < fastAnsList[j].a
				}
			})
			log.Printf("Slow verbose answers:\n%v\n", slowAnsList)
			log.Printf("Fast verbose answers:\n%v\n", fastAnsList)
			t.Fatalf("Slow (%v) and Fast (%v) solvers return different answers on N = %v", slowAns, fastAns, N)
		}
	}
}
