package sort

// 计数排序
func CountingSort(data []int) {
	if len(data) <= 1 {
		return
	}

	// 寻找最大值
	max := data[0]
	for _, v := range data {
		if max < v {
			max = v
		}
	}

	// 记录数值出现次数
	count := make([]int, max+1)
	for _, v := range data {
		count[v]++
	}

	// count数组累加
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	temp := make([]int, len(data))
	for _, v := range data {
		temp[count[v]-1] = v
		count[v]--
	}

	copy(data[:], temp[:])
}
