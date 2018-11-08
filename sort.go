package gutil

/*
	IntInsertionSortAsc:插入排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func IntInsertionSortAsc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// 小于前一个
		if arr[i] < arr[i-1] {
			index := i - 1
			temp := arr[i]
			// 为元素找到合适的位置
			for index >= 0 && arr[index] > temp {
				arr[index+1] = arr[index]
				index--
			}
			arr[index+1] = temp
		}
	}
	return arr
}

/*
	Int32InsertionSortAsc:插入排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func Int32InsertionSortAsc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// 小于前一个
		if arr[i] < arr[i-1] {
			index := i - 1
			temp := arr[i]
			// 为元素找到合适的位置
			for index >= 0 && arr[index] > temp {
				arr[index+1] = arr[index]
				index--
			}
			arr[index+1] = temp
		}
	}
	return arr
}

/*
	Int64InsertionSortAsc:插入排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func Int64InsertionSortAsc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// 小于前一个
		if arr[i] < arr[i-1] {
			index := i - 1
			temp := arr[i]
			// 为元素找到合适的位置
			for index >= 0 && arr[index] > temp {
				arr[index+1] = arr[index]
				index--
			}
			arr[index+1] = temp
		}
	}
	return arr
}

/*
	IntInsertionSortDesc:插入排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func IntInsertionSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// 小于前一个
		if arr[i] > arr[i-1] {
			index := i - 1
			temp := arr[i]
			// 为元素找到合适的位置
			for index >= 0 && arr[index] < temp {
				arr[index+1] = arr[index]
				index--
			}
			arr[index+1] = temp
		}
	}
	return arr
}

/*
	Int32InsertionSortDesc:插入排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func Int32InsertionSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// 小于前一个
		if arr[i] > arr[i-1] {
			index := i - 1
			temp := arr[i]
			// 为元素找到合适的位置
			for index >= 0 && arr[index] < temp {
				arr[index+1] = arr[index]
				index--
			}
			arr[index+1] = temp
		}
	}
	return arr
}

/*
	Int64InsertionSortDesc:插入排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func Int64InsertionSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// 小于前一个
		if arr[i] > arr[i-1] {
			index := i - 1
			temp := arr[i]
			// 为元素找到合适的位置
			for index >= 0 && arr[index] < temp {
				arr[index+1] = arr[index]
				index--
			}
			arr[index+1] = temp
		}
	}
	return arr
}
