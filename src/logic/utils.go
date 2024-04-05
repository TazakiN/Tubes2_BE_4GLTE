package logic

import (
	"fmt"
)

type Node struct {
	// visited    bool
	link       string
	neighbours []*Node
	distance   int
	parent     *Node
}

func newNode(link string) *Node {
	return &Node{
		link: link,
		// visited:    false,
		neighbours: []*Node{},
		distance:   0,
		parent:     nil,
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("[%s]", getPageTitle(n.link))
}
