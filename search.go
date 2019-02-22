package gutil

/*
	IntBinarySearch 二分查找
	input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 1
	output: 1 or -1 (if key !exists result=-1)
*/
func IntBinarySearch(array []int, key int) int {
	low, hight := 0, len(array)-1
	for low <= hight {
		m := (low + hight) >> 1
		if array[m] < key {
			low = m + 1
		} else if array[m] > key {
			hight = m - 1
		} else {
			// 相等
			return m
		}
	}
	return -1
}

/*
	Int32BinarySearch 二分查找
	input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 1
	output: 1 or -1 (if key !exists result=-1)
*/
func Int32BinarySearch(array []int32, key int32) int {
	low, hight := 0, len(array)-1
	for low <= hight {
		m := (low + hight) >> 1
		if array[m] < key {
			low = m + 1
		} else if array[m] > key {
			hight = m - 1
		} else {
			// 相等
			return m
		}
	}
	return -1
}

/*
	Int64BinarySearch 二分查找
	input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 1
	output: 1 or -1 (if key !exists result=-1)
*/
func Int64BinarySearch(array []int64, key int64) int {
	low, hight := 0, len(array)-1
	for low <= hight {
		m := (low + hight) >> 1
		if array[m] < key {
			low = m + 1
		} else if array[m] > key {
			hight = m - 1
		} else {
			// 相等
			return m
		}
	}
	return -1
}

// Int32InSlice 判断int32是否在切片里面
func Int32InSlice(needle int32, haystack []int32) bool {
	for _, b := range haystack {
		if b == needle {
			return true
		}
	}
	return false
}

// StringInSlice 判断字符是否在切片里面
func StringInSlice(needle string, haystack []string) bool {
	for _, b := range haystack {
		if b == needle {
			return true
		}
	}
	return false
}

// Int64InSlice 判断int64是否在切片里面
func Int64InSlice(needle int64, haystack []int64) bool {
	for _, b := range haystack {
		if b == needle {
			return true
		}
	}
	return false
}

// IntInSlice 判断int是否在切片里面
func IntInSlice(needle int, haystack []int) bool {
	for _, b := range haystack {
		if b == needle {
			return true
		}
	}
	return false
}

// InterfaceInSlice 判断interface是否在切片里面
func InterfaceInSlice(needle interface{}, haystack []interface{}) bool {
	for _, b := range haystack {
		if b == needle {
			return true
		}
	}
	return false
}
