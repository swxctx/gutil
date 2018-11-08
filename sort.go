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
func Int32InsertionSortAsc(arr []int32) []int32 {
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
func Int64InsertionSortAsc(arr []int64) []int64 {
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
func Int32InsertionSortDesc(arr []int32) []int32 {
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
func Int64InsertionSortDesc(arr []int64) []int64 {
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
	IntBubbleSortAsc 冒泡排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func IntBubbleSortAsc(values []int) []int {
	flag := true
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		// 使用flag标记，减少循环次数(当第一次已经是正序就不再进行循环了)
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
	return values
}

/*
	IntBubbleSortAsc 冒泡排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func Int32BubbleSortAsc(values []int32) []int32 {
	flag := true
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		// 使用flag标记，减少循环次数(当第一次已经是正序就不再进行循环了)
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
	return values
}

/*
	Int64BubbleSortAsc 冒泡排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func Int64BubbleSortAsc(values []int64) []int64 {
	flag := true
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		// 使用flag标记，减少循环次数(当第一次已经是正序就不再进行循环了)
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
	return values
}

/*
	IntBubbleSortDesc 冒泡排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func IntBubbleSortDesc(values []int) []int {
	flag := true
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		// 使用flag标记，减少循环次数(当第一次已经是正序就不再进行循环了)
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] < values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
	return values
}

/*
	Int32BubbleSortDesc 冒泡排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func Int32BubbleSortDesc(values []int32) []int32 {
	flag := true
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		// 使用flag标记，减少循环次数(当第一次已经是正序就不再进行循环了)
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] < values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
	return values
}

/*
	Int64BubbleSortDesc 冒泡排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func Int64BubbleSortDesc(values []int64) []int64 {
	flag := true
	vLen := len(values)
	for i := 0; i < vLen-1; i++ {
		// 使用flag标记，减少循环次数(当第一次已经是正序就不再进行循环了)
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values[j] < values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
	return values
}

/*
	IntSectionSortAsc 选择排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func IntSectionSortAsc(values []int) []int {
	length := len(values)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if values[j] < values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}

/*
	Int32SectionSortAsc 选择排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func Int32SectionSortAsc(values []int32) []int32 {
	length := len(values)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if values[j] < values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}

/*
	Int64SectionSortAsc 选择排序
	input: [95 45 15 78 84 51 24 12]
	output: [12 15 24 45 51 78 84 95]
*/
func Int64SectionSortAsc(values []int64) []int64 {
	length := len(values)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if values[j] < values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}

/*
	IntSectionSortDesc 选择排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func IntSectionSortDesc(values []int) []int {
	length := len(values)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if values[j] > values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}

/*
	Int32SectionSortDesc 选择排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func Int32SectionSortDesc(values []int32) []int32 {
	length := len(values)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if values[j] > values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}

/*
	IntSectionSortDesc 选择排序
	input: [95 45 15 78 84 51 24 12]
	output: [95 84 78 51 45 24 15 12]
*/
func Int64SectionSortDesc(values []int64) []int64 {
	length := len(values)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if values[j] > values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}
	return values
}
