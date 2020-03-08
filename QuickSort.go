package main

import "fmt"

func main() {
	a := []int{15, 5, 6, 3, 4}
	fmt.Println("unsorted array: ", a)
	n := len(a)
	lb := 0
	ub := n - 1
	QuickSort(a, lb, ub)
	fmt.Println("sorted array: ", a)
}

func QuickSort(a []int, low int, high int) {

	if low < high {
		pivot := Partition(a, low, high)
		QuickSort(a, low, pivot)
		QuickSort(a, pivot+1, high)
	}

}

func Partition(a []int, low, high int) int {
	pivot := a[low]
	i := low
	j := high

	for i < j {

		for a[i] <= pivot && i < high {
			i++
		}
		for a[j] > pivot && j > low {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}

	}
	a[low] = a[j]
	a[j] = pivot

	return j
}
