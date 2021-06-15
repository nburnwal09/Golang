package main

import "fmt"

const max int = 5

type stack struct {
	st  [max]int
	top int
}

func (s *stack) push(val int) {
	if s.top >= max {
		fmt.Println("stack full")
		return
	}
	s.top = s.top + 1
	s.st[s.top] = val
}

func (s *stack) pop() {
	if s.top == -1 {
		fmt.Println("Stack empty")
		return
	}
	s.top--
}
func (s stack) print() {
	if s.top == -1 {
		fmt.Println("Stack is empty")
		return
	}
	for i := 0; i <= s.top; i++ {
		fmt.Print(s.st[i], "   ")
	}
	println()
}
func main() {
	s := stack{top: -1}
	for true {
		fmt.Println("Enter 1-Push 2-Pop 3-Print")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var val int
			fmt.Println("Enter value:-")
			fmt.Scanln(&val)
			s.push(val)
		case 2:
			s.pop()
		case 3:
			fmt.Println("Stack is:- ")
			s.print()
		default:
			fmt.Println("Invalid choice")
		}
	}
}
