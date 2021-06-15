package main

import "fmt"

type queue struct {
	arr  []int
	top  int
	rear int
}

func (q *queue) enqueue(val int) {
	if q.top == -1 {
		q.top++
	}
	q.rear++
	q.arr = append(q.arr, val)
}

func (q *queue) dequeue() {
	if q.top == -1 {
		fmt.Println("queue empty")
	} else if q.top == q.rear {
		q.rear = -1
		q.top = -1
	} else {
		q.top++
	}
}
func (q queue) print() {
	fmt.Println("top=", q.top, "  q.rear=", q.rear)
	if (q.top == q.rear+1) || (q.top == -1) {
		fmt.Println("Empty queue")
		return
	}
	for i := q.rear; i >= q.top && i >= 0; i-- {
		fmt.Print(q.arr[i], "   ")
	}
	println()
}
func main() {
	q := queue{top: -1, rear: -1}
	for true {
		fmt.Println("Enter 1-Enqueue 2-Dequeue 3-Print")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var val int
			fmt.Println("Enter value:-")
			fmt.Scanln(&val)
			q.enqueue(val)
		case 2:
			q.dequeue()
		case 3:
			fmt.Println("Queue is:- ")
			q.print()
		default:
			fmt.Println("Invalid choice")
		}
	}

}
