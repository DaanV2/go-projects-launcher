package slicesx

// Filter removes any item that doesn't match the predicate
func Filter[S ~[]E, E any](items S, predicate func(E) bool) S {
	result := make(S, 0, len(items))

	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
