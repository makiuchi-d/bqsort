package bqsort_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/makiuchi-d/bqsort"
)

const CNT = 10000
const MAXKEY = 500

func genInts(n int) []Ints {
	ints := make([]Ints, n)
	for i := range ints {
		ints[i] = make(Ints, CNT)
		for j := range ints[i] {
			ints[i][j] = rand.Intn(MAXKEY)
		}
	}
	return ints
}

func BenchmarkSort(b *testing.B) {
	ints := genInts(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(ints[i])
	}
}

func BenchmarkBQSort(b *testing.B) {
	ints := genInts(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bqsort.Sort(ints[i], MAXKEY)
	}
}
