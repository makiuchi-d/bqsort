package bqsort_test

import (
	"math"
	"math/rand"
	"reflect"
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

func TestReverse(t *testing.T) {
	data := Ints{3, 5, 2, 1, 4}
	rev := bqsort.Reverse(data)

	bqsort.Sort(rev, 5)

	exp := Ints{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("reverse ints are not sorted: %v", data)
	}

	rr := bqsort.Reverse(rev)
	if reflect.ValueOf(data).Pointer() != reflect.ValueOf(rr).Pointer() {
		t.Fatalf("reverse again must return the original object: %p %p", data, rr)
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

func TestInt8s(t *testing.T) {
	data := []int8{127, -15, 10, -1, 0}
	exp := []int8{-15, -1, 0, 10, 127}
	bqsort.Int8s(data)
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("data is not sorted: %x, wants %x", data, exp)
	}
}

func TestUint8s(t *testing.T) {
	data := []uint8{127, 15, 10, 1, 0}
	exp := []uint8{0, 1, 10, 15, 127}
	bqsort.Uint8s(data)
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("data is not sorted: %x, wants %x", data, exp)
	}
}

func TestInt16s(t *testing.T) {
	data := []int16{127, -15, 10, -1, 1024, 0, -500}
	exp := []int16{-500, -15, -1, 0, 10, 127, 1024}
	bqsort.Int16s(data)
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("data is not sorted: %x, wants %x", data, exp)
	}
}

func TestUint16s(t *testing.T) {
	data := []uint16{127, 15, 10, 1, 1024, 0, 500}
	exp := []uint16{0, 1, 10, 15, 127, 500, 1024}
	bqsort.Uint16s(data)
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("data is not sorted: %x, wants %x", data, exp)
	}
}

func TestInt32s(t *testing.T) {
	data := []int32{65536, 127, -15, 10, -1, 1024, 0, -500}
	exp := []int32{-500, -15, -1, 0, 10, 127, 1024, 65536}
	bqsort.Int32s(data)
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("data is not sorted: %x, wants %x", data, exp)
	}
}

func TestUint32s(t *testing.T) {
	data := []uint32{65536, 127, 15, 10, 1, 1024, 0, 500}
	exp := []uint32{0, 1, 10, 15, 127, 500, 1024, 65536}
	bqsort.Uint32s(data)
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("data is not sorted: %x, wants %x", data, exp)
	}
}

func TestFloats32(t *testing.T) {
	data := []float32{
		3.5, float32(math.Inf(1)), -10.0, math.MaxFloat32, -math.MaxFloat32, float32(math.Inf(-1)), 0, float32(math.NaN()),
	}
	exp := []float32{
		float32(math.NaN()), float32(math.Inf(-1)), -math.MaxFloat32, -10.0, 0, 3.5, math.MaxFloat32, float32(math.Inf(1)),
	}
	bqsort.Float32s(data)
	if !math.IsNaN(float64(data[0])) {
		t.Fatalf("NaN must be placed on top: %v", data)
	}
	if !reflect.DeepEqual(data[1:], exp[1:]) { // skip NaN
		t.Fatalf("data is not sorted: %v, wants %v", data, exp)
	}
}
