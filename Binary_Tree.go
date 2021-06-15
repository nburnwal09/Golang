package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type bTree struct {
	root *Node
}

func (bt *bTree) insert(val int) {
	if bt.root == nil {
		bt.root = &Node{data: val}
	} else {
		bt.root.insert(val)
	}
}
func (n *Node) insert(val int) {
	if val <= n.data {
		if n.left == nil {
			n.left = &Node{data: val}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: val}
		} else {
			n.right.insert(val)
		}
	}
}
func (n *Node) preOrder() {
	if n == nil {
		return
	}
	fmt.Print(n.data, "  ")
	n.left.preOrder()
	n.right.preOrder()
}
func (root *Node) height() int { //12

	if root == nil {
		return 0
	}
	lHeight := root.left.height() //1 //root=2(1)
	rHeight := root.right.height() //2  //root=14(2)
	if lHeight > rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1 //3
	}
}

func (root *Node) printLevel(l int) {
	if root == nil {
		return
	}
	if l == 1 {
		fmt.Print(root.data, "  ")
		return
	} else {
		root.left.printLevel(l - 1)
		root.right.printLevel(l - 1)
	}
}
func (b *bTree) levelOrder() {
	h := b.root.height()
	for i := 1; i <= h; i++ {
		b.root.printLevel(i)//i=1(12), i=2(2,14), i=3
	}
}
func main() {
	var b1 bTree
	b1.insert(12)
	b1.insert(2)
	b1.insert(13)
	b1.insert(4)
	b1.insert(15)
	b1.insert(16)
	b1.insert(14)
	b1.insert(25)
	fmt.Println("Pre Order:")
	b1.root.preOrder()
	fmt.Println(b1.root.height(), "\n")
	b1.levelOrder()

	// b1.root = &Node{1, nil, nil}
	// b1.root.left = &Node{2, nil, nil}
	// b1.root.right = &Node{3, nil, nil}
	// b1.root.left.left = &Node{4, nil, nil}
	// b1.root.left.right = &Node{5, nil, nil}
	// // b1.insertData(1)
	// // b1.insertData(2)
	// // b1.insertData(3)
	// // b1.insertData(4)
	// fmt.Println(b1.root.data)
	// fmt.Println(b1.root.left.data, "  ", b1.root.right.data)
	// fmt.Println(b1.root.left.left.data, "  ", b1.root.left.right.data)

}


// type slNode []Node
// snode := []slNode{}
// temp := b.root
// for temp != nil {
// 	fmt.Println(temp.data)
// 	if temp.left != nil {
// 		snode = append(snode, temp.left)
// 	} else if temp.right != nil {
// 		snode = append(snode, temp.right)
// 	}
// 	temp = snode[len(snode)-1]
// 	snode = snode[:len(snode)-1]
// }