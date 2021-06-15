package main

import "fmt"

func merge(x []int, y []int) []int {
	ans := []int{}
	lx := len(x)
	ly := len(y)
	i, j := 0, 0
	for i < lx && j < ly {
		if x[i] <= y[j] {
			ans = append(ans, x[i])
			i++
		} else {
			ans = append(ans, y[j])
			j++
		}
	}
	for i < lx {
		ans = append(ans, x[i])
		i++
	}
	for j < ly {
		ans = append(ans, y[j])
		j++
	}
	return ans
}
func main() {
	a := []int{1, 2, 8, 30, 60, 90}
	b := []int{3, 5, 6, 14, 36, 50, 100}
	fmt.Println(merge(a, b))
}
