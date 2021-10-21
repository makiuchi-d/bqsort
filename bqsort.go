// Package bqsort provides in-space binary MSD radix sort (a.k.a. binary quick sort).
// The cost of this algorithm is O(n*k), where n is the data count and
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
