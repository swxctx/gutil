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
