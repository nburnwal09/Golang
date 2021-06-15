package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	var j int
	for i := 1; i < len(a); i++ {
		key := a[i]
		j = i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j = j - 1
		}
		fmt.Println("j=", j)
		a[j+1] = key
	}
	fmt.Println(a)
}
