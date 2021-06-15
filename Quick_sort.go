package main

import "fmt"

func partition(a []int, low int, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}
func quickSort(a []int, low int, high int) {
	if low < high {
		pi := partition(a, low, high)
		quickSort(a, 0, pi-1)
		quickSort(a, pi+1, high)
	}
}
func main() {
	a := []int{52, 6, 32, 98, 12, 25, 56, -5, -6, -36}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
