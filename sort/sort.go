package sort

// InsertSort 插入排序
func InsertSort(data []int) []int {
	if data == nil || len(data) < 2 {
		return data
	}
	for j := 1; j < len(data); j++ {
		temp := data[j]
		i := j - 1
		for ; i >= 0 && data[i] > temp; i-- {
			data[i+1] = data[i]
		}
		data[i+1] = temp
	}
	return data
}
