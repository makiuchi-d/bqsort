# bqsort
Binary quicksort for Go.

[![Test](https://github.com/makiuchi-d/bqsort/actions/workflows/test.yml/badge.svg)](https://github.com/makiuchi-d/bqsort/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/makiuchi-d/bqsort/branch/main/graph/badge.svg?token=8PYEVZ35U9)](https://codecov.io/gh/makiuchi-d/bqsort)

Binary quicksort is a in-space binary MSD radix sort.
This is not a stable sort.
In the worst case, this sort takes **O(n*k)** evaluation and uses **O(k)** space complexity,
where *n* is the element count and *k* is the number of bits in the max key value.

This algorithm is advantageous when the number of keys is limited and the number of elements is large.


## Go Docs

See: [pkg.go.dev](https://pkg.go.dev/github.com/makiuchi-d/bqsort)
