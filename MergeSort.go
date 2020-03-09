package main
// https://play.golang.org/p/dS6L-u8MyKA
import "fmt"

func main() {

	a := []int{10, 5, 3, 4, 7}
	fmt.Println("unsorted array: ", a)

	n := len(a)
	lb, ub := 0, n-1
	MergeSort(a, lb, ub)
	fmt.Println("sorted array: ", a)
}

func MergeSort(a []int, lb, ub int) {
	fmt.Println("lb < ub : lb , ub : ", lb, ub)
	if lb < ub {
		mid := (lb + ub) / 2
		fmt.Println("divide from lb to mid : lb , mid : ", lb, mid)
		fmt.Println("array:", a[lb:mid+1])
		MergeSort(a, lb, mid)
		fmt.Println("divide from mid+1 to ub : mid+1 , ub : ", mid+1, ub)
		fmt.Println("array:", a[mid+1:ub+1])
		MergeSort(a, mid+1, ub)
		fmt.Println("merging from lb  mid ub : lb  mid ub: ", lb, mid, ub)
		Merge(a, lb, mid, ub)
	}

}

func Merge(a []int, lb, mid, ub int) {
	i := lb
	j := mid + 1
	b := []int{}
	for i <= mid && j <= ub {
		fmt.Println("first half:", a[i:mid+1])
		fmt.Println("second half: ", a[mid+1:ub+1])
		if a[i] <= a[j] {
			fmt.Println("a[i]<a[j]: so appending a[i] to b", a[i], a[j])
			b = append(b, a[i])
			fmt.Println("b after appending: ", b)
			i++
			fmt.Println("incrementing I: ", i)
		} else {
			fmt.Println("a[i]>a[j]: so appending a[j] to b", a[i], a[j])
			b = append(b, a[j])
			fmt.Println("b after appending: ", b)
			j++
			fmt.Println("incrementing j: ", j)
		}
		fmt.Println("i ,mid,j,ub in for loop", i, mid, j, ub)
	}

	if i > mid {
		for j <= ub {
			fmt.Println("appending all remaining from second half")
			b = append(b, a[j])
			fmt.Println("b after appending: ", b)
			j++
		}
	} else {
		for i <= mid {
			fmt.Println("appending all remaining from second half")
			b = append(b, a[i])
			fmt.Println("b after appending: ", b)
			i++
		}
	}

	for k := range b {
		a[k] = b[k]
	}
}
