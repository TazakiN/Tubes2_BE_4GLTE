package logic

import (
	"fmt"
)

type Node struct {
	// visited    bool
	link       string
	title      string
	neighbours []*Node
	parent     *Node
}

func newNode(link string, title string) *Node {
	return &Node{
		// visited:    false,
		link:       link,
		title:      title,
		neighbours: []*Node{},
		parent:     nil,
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("[%s]", getPageTitle(n.link))
}

func reverse(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}
