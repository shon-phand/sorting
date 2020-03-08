package main

import "fmt"

/*

best time complexity:  o(n)
worst case : o(n`2)
*/

func main() {

	arr := []int{16, 14, 5, 6, 8}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println("sorted arr: ", arr)
}

func BubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ { // this is for passed
		k := 0
		flag := 0
		for j := 0; j < n-1-i; j++ { /// j-1-i reduced number of iteration // this is for iterations
			k++
			if arr[j+1] < arr[j] {
				flag = 1
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Printf("iterations %d", k)
		fmt.Println(" passs ", i, ":", arr)

		if flag == 0 { // this flag is for optimized bubble sort
			break // it reduces the passes
		}

	}
}
