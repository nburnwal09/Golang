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
	text = strings.TrimSpace(text)
	fmt.Printf("%T", text)
	t := strings.Split(text, " ")
	fmt.Printf("%v", t)
	a := []int{}
	for _, v := range t {
		nn, _ := strconv.Atoi(v)
		a = append(a, nn)
	}
	//fmt.Print(a)
	// a := []int{}
	// var l int
	// fmt.Scanln(&l)
	// for i := 0; i < l; i++ {
	// 	var n int
	// 	fmt.Scanln(&n)
	// 	a = append(a, n)
	// }
	for i := 0; i < len(a)-1; i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}

		}
	}
	fmt.Println(a)
}
