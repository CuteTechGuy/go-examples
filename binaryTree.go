package main

import (
	"fmt"
)

var count int = 0

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
}

// Insert

func (n *Node) Insert(k int, v string) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k, Value: v}
		} else {
			n.Right.Insert(k, v)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k, Value: v}
		} else {
			n.Left.Insert(k, v)
		}
	}
}

// Search

func (n *Node) Search(k int) (bool, string) {
	count++
	if n == nil {
		return false, ""
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true, n.Value
}

func main() {
	tree := &Node{Key: 100, Value: "Birinci"}
	tree.Insert(50, "İkinci")
	tree.Insert(105, "Üçüncü")
	tree.Insert(35, "Dördüncü")
	tree.Insert(22, "Beşinci")
	tree.Insert(15, "Altıncı")
	tree.Insert(123, "Yedinci")
	tree.Insert(125, "Sekizinci")

	fmt.Println(tree)
	found, value := tree.Search(14)
	fmt.Printf("Found: %v | Value: %v | Search Count: %v", found, value, count)

}
