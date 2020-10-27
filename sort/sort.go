package sort

// 冒泡排序
func Bubble(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}

	for i := len - 1; i >= 0; i-- {
		stop := true
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				stop = false
			}
		}
		if stop {
			break
		}
	}

	return arr
}

// 插入排序
func Insertion(arr []int) []int {
	ln := len(arr)
	if ln <= 1 {
		return arr
	}

	sorted := arr[:1]
	unsorted := arr[1:]

	unLen := len(unsorted)
	for i := 0; i < unLen; i++ {
		tgt := unsorted[i]
		sorted = append(sorted, tgt)
		for j := 0; j < len(sorted); j++ {
			if sorted[j] > tgt {
				last2index(sorted, j)
				break
			}
		}
	}

	return sorted
}

// 选择排序
func Selection(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	for i := 0; i < len; i++ {
		minIndex := min(arr[i:]) + i
		swap(arr, i, minIndex)
	}
	return arr
}

// 归并排序
func Merge(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}

	divide(arr)
	return arr
}

// 分解合并的单元操作
func divide(arr []int) {
	l := len(arr)
	i := l / 2
	if l > 2 {
		divide(arr[:i])
		divide(arr[i:])
	}

	merge(arr[:i], arr[i:])
}

// 合并两个数组
func merge(arr1, arr2 []int) {

	// 比较两边头部值 小者放前
	l1 := len(arr1)
	l2 := len(arr2)

	var arr []int
	j1 := 0
	j2 := 0
	for {
		if j1 > l1-1 && j2 > l2-1 {
			break
		}
		k1 := j1
		if j1 < l1 && (j2 == l2 || arr1[j1] < arr2[j2]) {
			arr = append(arr, arr1[j1])
			j1++
		}

		if j2 < l2 && (k1 == l1 || arr1[k1] >= arr2[j2]) {
			arr = append(arr, arr2[j2])
			j2++
		}
	}

	index := (l1 + l2) / 2
	copy(arr1, arr[:index])
	copy(arr2, arr[index:])
}

// 快速排序
func Quick(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	piv := arr[l-1]
	i := 0
	for j := 0; j < l-1; j++ {
		if arr[j] < piv {
			swap(arr, i, j)
			i++
		}
	}
	swap(arr, l-1, i)

	// 最后一个替换位
	lastI := i - 1
	if lastI < 0 {
		lastI = 0
	}

	if i > 0 {
		Quick(arr[:i])
	}
	if i < l-1 {
		Quick(arr[i+1:])
	}

	return arr
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

// 数组最大值的索引
func min(arr []int) int {
	min := INT_MAX
	index := 0
	for i, v := range arr {
		if v < min {
			min = v
			index = i
		}
	}
	return index
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// 末位元素移到指定位置
func last2index(arr []int, i int) {
	l := len(arr)
	tmp := append([]int{arr[l-1]}, arr[i:l-1]...)
	copy(arr[i:], tmp)
}
