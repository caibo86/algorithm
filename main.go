package main

import (
	"fmt"
	"github.com/caibo86/algorithm/sort"
)

func main() {
	ret := []int{0, 1, 3, 4, 10, 12, 15, 18, 22, 31, 33, 45, 89, 92, 199, 200}
	data := sort.InsertSort(getData())
	if compare(ret, data) {
		fmt.Println("插入排序成功")
	} else {
		fmt.Println("插入排序失败")
	}
}

func getData() []int {
	data := []int{22, 33, 1, 4, 15, 0, 200, 199, 45, 31, 12, 18, 92, 89, 10, 3}
	return data
}

func compare(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
