package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubble(a []int, l int) {
	if l == 1 {
		return
	}
	for i := 1; i <= l-1; i++ {
		if a[i] < a[i-1] {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}
	bubble(a, l-1)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	text = text[:len(text)-2]
	t := strings.Split(text, " ")
	a := []int{}
	for _, n := range t {
		nn, _ := strconv.Atoi(n)
		a = append(a, nn)
	}
	bubble(a, len(a))
	fmt.Println(a)
}
