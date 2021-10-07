package sort

import "testing"

func TestCountingSort(t *testing.T) {
	data := []int{1}
	CountingSort(data)
	t.Log(data)

	data1 := []int{1, 0}
	CountingSort(data1)
	t.Log(data1)

	data2 := []int{1, 0, 2}
	CountingSort(data2)
	t.Log(data2)

	data3 := []int{1, 0, 1, 1, 1, 0, 0}
	CountingSort(data3)
	t.Log(data3)
}
