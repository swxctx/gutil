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
