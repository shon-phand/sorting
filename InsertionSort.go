package main

import (
	"fmt"
)

/*

bestcase time complexity o(n)
worst case time complexity o(n`2)

*/

func main() {
	a := []int{10, 3, 5, 9, 7}
	fmt.Println("unsortd array", a)
	InsertionSort(a)
	fmt.Println("unsortd array", a)

}

func InsertionSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ { // for unsorted array traversal

		j := i - 1                  //  placing pointer at end of sorted array
		temp := a[i]                // taking backup of element needs to insertion
		for j >= 0 && a[j] > temp { // compare in a sorted array for correct position
			a[j+1] = a[j] // moving the element if the element in a sorted array is greater than element in sorted array
			j--
		}
		a[j+1] = temp // placing element in correct position
	}
}
