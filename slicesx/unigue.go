package slicesx

import "slices"

func Unique[S ~[]E, E comparable](items S) S {
	result := make(S, 0, len(items))

	for _, item := range items {
		if slices.Contains(result, item) {
			continue
		}
		result = append(result, item)
	}

	return result
}
