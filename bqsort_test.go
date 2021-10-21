package bqsort_test

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/makiuchi-d/bqsort"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Ints []int

func (x Ints) Len() int           { return len(x) }
func (x Ints) Less(i, j int) bool { return x[i] < x[j] }
func (x Ints) Key(i int) uint64   { return uint64(x[i]) }
func (x Ints) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func TestSort(t *testing.T) {
	data := make([]int, 5)
	max := 0
	for i := range data {
		data[i] = rand.Int()
		if max < data[i] {
			max = data[i]
		}
	}

	bqsort.Sort(Ints(data), uint64(max))

	if !sort.IntsAreSorted(data) {
		t.Fatalf("ints are not sorted: %x", data)
	}
}

func TestSlice(t *testing.T) {
	data := make([]uint8, 10)
	for i := range data {
		data[i] = uint8(rand.Intn(math.MaxUint8))
	}

	bqsort.Slice(data, math.MaxUint8, func(i int) uint64 { return uint64(math.MaxUint8 - data[i]) })

	if !sort.SliceIsSorted(data, func(i, j int) bool { return data[i] > data[j] }) {
		t.Fatalf("slice is not sorted: %x", data)
	}
}
