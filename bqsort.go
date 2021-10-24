// Package bqsort provides in-space binary MSD radix sort (a.k.a. binary quick sort).
// The cost of this algorithm is O(n*k), where n is the element count and
// k is the number of bits in the max key value.
package bqsort

import (
	"math"
	"reflect"
)

// An implementation of Interface can be sorted by the routines in this package.
// The methods refer to elements of the underlying collection by integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Key returns sort-key value of the element with index i.
	// The smaller key elements come before larger keys.
	Key(i int) uint64

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type keySwapper interface {
	Key(i int) uint64
	Swap(i, j int)
}

type keySwap struct {
	key  func(i int) uint64
	swap func(i, j int)
}

func (ks keySwap) Key(i int) uint64 { return ks.key(i) }
func (ks keySwap) Swap(i, j int)    { ks.swap(i, j) }

// msb returns most significant bit (MSB).
func msb(val uint64) uint64 {
	bit := uint64((math.MaxUint64 + 1) >> 1)
	for val < bit {
		bit >>= 1
	}
	return bit
}

func sort(data keySwapper, a, b int, bit uint64) {
	ma := a
	mb := b

	for ma < mb {
		for ma < mb && data.Key(ma)&bit == 0 {
			ma++
		}
		mb--
		for ma < mb && data.Key(mb)&bit != 0 {
			mb--
		}
		if ma < mb {
			data.Swap(ma, mb)
			ma++
		}
	}

	bit >>= 1
	if bit == 0 {
		return
	}
	if a < ma-1 {
		sort(data, a, ma, bit)
	}
	if ma < b-1 {
		sort(data, ma, b, bit)
	}
}

// Sort sorts data.
// The data.Key() must be less than or equal to maxkey.
// The cost is O(n*k), where n is the data.Len() and k is the number of bits in the maxkey.
func Sort(data Interface, maxkey uint64) {
	sort(data, 0, data.Len(), msb(maxkey))
}

// Slice sorts a slice.
// The key(i) returns the sort key of the slice[i] which must be less than or equal to maxkey.
// The cost is O(n*k), where n is the len(x) and k is the number of bits in the maxkey.
func Slice(slice interface{}, maxkey uint64, key func(i int) uint64) {
	len := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	sort(keySwap{key, swap}, 0, len, msb(maxkey))
}

// Int8Slice implements Interface for []int8.
type Int8Slice []int8

func (x Int8Slice) Len() int         { return len(x) }
func (x Int8Slice) Key(i int) uint64 { return uint64(x[i]) + (-math.MinInt8) }
func (x Int8Slice) Swap(i, j int)    { x[i], x[j] = x[j], x[i] }

// Int8s sorts []int8.
// The cost is O(n*8).
func Int8s(x []int8) {
	Sort(Int8Slice(x), msb(math.MaxInt8+(-math.MinInt8)))
}

// Uint8Slice implements Interface for []uint8.
type Uint8Slice []uint8

func (x Uint8Slice) Len() int         { return len(x) }
func (x Uint8Slice) Key(i int) uint64 { return uint64(x[i]) }
func (x Uint8Slice) Swap(i, j int)    { x[i], x[j] = x[j], x[i] }

// Uint8s sorts []uint8.
// The cost is O(n*8).
func Uint8s(x []uint8) {
	Sort(Uint8Slice(x), msb(math.MaxUint8))
}

// Int16Slice implements Interface for []int16.
type Int16Slice []int16

func (x Int16Slice) Len() int         { return len(x) }
func (x Int16Slice) Key(i int) uint64 { return uint64(x[i]) + (-math.MinInt16) }
func (x Int16Slice) Swap(i, j int)    { x[i], x[j] = x[j], x[i] }

// Int16s sorts []int16.
// The cost is O(n*16).
func Int16s(x []int16) {
	Sort(Int16Slice(x), msb(math.MaxInt16+(-math.MinInt16)))
}

// Uint16Slice implements Interface for []uint16.
type Uint16Slice []uint16

func (x Uint16Slice) Len() int         { return len(x) }
func (x Uint16Slice) Key(i int) uint64 { return uint64(x[i]) }
func (x Uint16Slice) Swap(i, j int)    { x[i], x[j] = x[j], x[i] }

// Uint16s sorts []uint8.
// The cost is O(n*16).
func Uint16s(x []uint16) {
	Sort(Uint16Slice(x), msb(math.MaxUint16))
}

// Int32Slice implements Interface for []int32.
type Int32Slice []int32

func (x Int32Slice) Len() int         { return len(x) }
func (x Int32Slice) Key(i int) uint64 { return uint64(x[i]) + (-math.MinInt32) }
func (x Int32Slice) Swap(i, j int)    { x[i], x[j] = x[j], x[i] }

// Int32s sorts []int32.
// The cost is O(n*32).
func Int32s(x []int32) {
	Sort(Int32Slice(x), msb(math.MaxInt32+(-math.MinInt32)))
}

// Uint32Slice implements Interface for []uint32.
type Uint32Slice []uint32

func (x Uint32Slice) Len() int         { return len(x) }
func (x Uint32Slice) Key(i int) uint64 { return uint64(x[i]) }
func (x Uint32Slice) Swap(i, j int)    { x[i], x[j] = x[j], x[i] }

// Uint32s sorts []uint32.
// The cost is O(n*32).
func Uint32s(x []uint32) {
	Sort(Uint32Slice(x), msb(math.MaxUint32))
}
