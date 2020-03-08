package main

import (
	"fmt"
)

/*

The selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order) from unsorted part and putting it at the beginning. The algorithm maintains two subarrays in a given array.

1) The subarray which is already sorted.
2) Remaining subarray which is unsorted.

In every iteration of selection sort, the minimum element (considering ascending order) from the unsorted subarray is picked and moved to the sorted subarray.

time complexity

best case = o(n'2)
worst case = o(n'2)

*/

func main() {
	fmt.Println("Selection sort")
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("arr to sorted:  ", arr)
	//SelectionSort(arr)
	n := len(arr)
	fmt.Println(n)
	selectionsort(arr)
	fmt.Println("arr to sorted:  ", arr)
}

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
