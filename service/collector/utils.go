package collector

import (
	"strconv"
	"strings"
)

func SliceInt64ToStr(values []int64, sep string) string {
	var valuesText []string
	for i := range values {
		text := strconv.FormatInt(values[i], 10)
		valuesText = append(valuesText, text)
	}

	// Join our string slice.
	result := strings.Join(valuesText, sep)
	return result

}

func chunkSlice(slice []int64, chunkSize int) [][]int64 {
	var chunks [][]int64
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func difference(a, b []int64) []int64 {
	mb := make(map[int64]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int64
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
func findNewProducts(a []int64, mb map[int64]bool) []int64 {
	var diff []int64
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
