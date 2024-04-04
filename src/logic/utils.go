package logic

import (
	"fmt"
)

type Node struct {
	link       string
	visited    bool
	neighbours []*Node
	distance   int
	parent     *Node
}

func newNode(link string) *Node {
	return &Node{
		link:       link,
		visited:    false,
		neighbours: []*Node{},
		distance:   0,
		parent:     nil,
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("[%s]", getPageTitle(n.link))
}

// func muatHasil(node *Node) []string {
// 	hasil := []string{}
// 	for node != nil {
// 		hasil = append(hasil, getPageTitle(node.link))
// 		node = node.parent
// 	}
// 	return hasil
// }
