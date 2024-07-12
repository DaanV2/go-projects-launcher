package slicesx

// Map applies a transformation function to each item in the slice
func Map[T any, U any](items []T, transform func(T) U) []U {
	result := make([]U, 0, len(items))

	for _, item := range items {
		result = append(result, transform(item))
	}

	return result
}
