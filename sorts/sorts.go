package sorts

import (
	"cmp"
	"slices"

	gers "github.com/PlayerR9/go-errors"
)

// CmpBucketSort returns a sorted slice of keys from the given table, sorted by the
// value returned by the given function. The returned slice is sorted in ascending order
// of the sorting keys.
//
// The function also supports duplicate keys. If two keys have the same sorting key,
// they will be sorted in the order they were added to the table.
//
// Parameters:
//   - table: a map from type to number of protections.
//   - fn: a function that takes a value from the table and returns an integer that
//     will be used as the sorting key.
//
// Returns:
//   - []K: a sorted slice of keys from the table.
//   - error: an error if the function is nil.
func CmpBucketSort[K cmp.Ordered, V any](table map[K]V, fn func(value V) int) ([]K, error) {
	if len(table) == 0 {
		return nil, nil
	} else if fn == nil {
		return nil, gers.NewErrNilParameter("fn")
	}

	buckets := make(map[int][]K)
	var indices []int

	for k, v := range table {
		weight := fn(v)

		var prev []K

		pos, ok := slices.BinarySearch(indices, weight)
		if !ok {
			prev = []K{k}
			indices = slices.Insert(indices, pos, weight)
		} else {
			prev, ok = buckets[weight]
			gers.AssertOk(ok, "buckets[%d]", weight)

			prev = append(prev, k)
		}

		buckets[weight] = prev
	}

	for _, bucket := range buckets {
		slices.Sort(bucket)
	}

	keys := make([]K, 0, len(table))

	for _, idx := range indices {
		vals, ok := buckets[idx]
		gers.AssertOk(ok, "buckets[%d]", idx)

		keys = append(keys, vals...)
	}

	return keys, nil
}
