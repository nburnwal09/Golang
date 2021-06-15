package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
	len  int
	// x:=10
	// var p *int//p poniter add(int)
	// p=&x
}

func (l *LinkedList) insertVal(val int) { //[1,next->(2,nil)]
	if l.head == nil { //                              temp
		l.head = &Node{data: val, next: nil}
		l.len++
	} else { //                             l.head       temp
		n := &Node{data: val, next: nil} // (1,next)---->(2,next)--->(3,nil)
		temp := l.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = n
		l.len++
	}
}
func (l *LinkedList) deleteNode(ind int) {
	if l.head == nil {
		fmt.Println("empty linked list")
		return
	}
	if ind >= l.len {
		fmt.Println("invalid index")
		return
	}
	//1 2 3 14 15 16
	temp := l.head        //temp=1
	if temp.next == nil { //if l.len==1{}
		temp = nil
		l.len = 0
		return
	}
	for i := 0; i < ind-1; i++ {
		temp = temp.next
	}
	dnode := temp.next
	//temp=3
	//t2=15
	t2 := l.head
	for i := 0; i < ind; i++ {
		t2 = t2.next
	}
	temp.next = t2.next
	dnode.next = nil
	l.len--
	return
}

func (l *LinkedList) reverse() {
	var prev, nxt *Node
	curr := l.head
	for curr != nil { //  12,(next) -> 2,(next) -> 14(next) -> 1(next) -> nil
		nxt = curr.next //
		curr.next = prev
		prev = curr
		curr = nxt
	}
	l.head = prev

	// var temp, nxt *Node
	// temp = l.head
	// for temp != nil {
	// 	temp, nxt, temp.next = temp.next, temp, nxt
	// }
	// l.head = nxt
}
func (l LinkedList) print() {
	temp := l.head
	if l.len == 0 {
		fmt.Println("Empty Linked List")
		return
	}
	for i := 0; i < l.len; i++ {
		fmt.Print(temp.data, "   ")
		temp = temp.next
	}
	println()

}
func main() { // 1->2->3->14->15->16
	var l1 LinkedList

	for true {
		fmt.Println("Enter 1-Insertion 2-Deletion 3-Print 4-Reverse and Print")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var val int
			fmt.Println("Enter value to insert:-")
			fmt.Scanln(&val)
			l1.insertVal(val)
		case 2:
			fmt.Println("Enter index to delete:-")
			var d int
			fmt.Scanln(&d)
			l1.deleteNode(d)
		case 3:
			fmt.Println("Linked List is:- ")
			l1.print()
		case 4:
			l1.reverse()
			l1.print()
		default:
			fmt.Println("Invalid choice")
		}
	}

}
