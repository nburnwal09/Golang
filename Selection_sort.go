package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-2]
	t := strings.Split(text, " ")
	a := []int{}
	for _, n := range t {
		num, _ := strconv.Atoi(n)
		a = append(a, num)
	}
	min := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}
	fmt.Println(a)
}
