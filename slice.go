package gutil

// AddToInt add a element to slice
func AddToInt(list []int, e int) []int {
	for _, l := range list {
		if l == e {
			return list
		}
	}
	return append(list, e)
}

// AddToInt32 add a element to slice
func AddToInt32(list []int32, e int32) []int32 {
	for _, l := range list {
		if l == e {
			return list
		}
	}
	return append(list, e)
}

// AddToInt64 add a element to slice
func AddToInt64(list []int64, e int64) []int64 {
	for _, l := range list {
		if l == e {
			return list
		}
	}
	return append(list, e)
}

// AddToFloat32 add a element to slice
func AddToFloat32(list []float32, e float32) []float32 {
	for _, l := range list {
		if l == e {
			return list
		}
	}
	return append(list, e)
}

// AddToFloat64 add a element to slice
func AddToFloat64(list []float64, e float64) []float64 {
	for _, l := range list {
		if l == e {
			return list
		}
	}
	return append(list, e)
}

// AddToString add a element to slice
func AddToString(list []string, e string) []string {
	for _, l := range list {
		if l == e {
			return list
		}
	}
	return append(list, e)
}

// AddToInterface add a element to slice
func AddToInterface(list []interface{}, e interface{}) []interface{} {
	for _, l := range list {
		if l == e {
			return list
		}
	}
	return append(list, e)
}
