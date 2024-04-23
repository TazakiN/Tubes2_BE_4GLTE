package logic

import (
	"fmt"
)

type Node struct {
	link   string
	title  string
	parent *Node
}

type NodeIDS struct {
	link   string
	title  string
	depth  int
	parent *NodeIDS
}

func newNode(link string, title string) *Node {
	return &Node{
		link:   link,
		title:  title,
		parent: nil,
	}
}

func newNodeIDS(link string, title string, depth int) *NodeIDS {
	return &NodeIDS{
		link:   link,
		title:  title,
		depth:  depth,
		parent: nil,
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

func getPath(node *Node) []string {
	path := []string{}
	for node != nil {
		path = append(path, node.link)
		node = node.parent
	}
	return reverse(path)
}

func getPathIDS(node *NodeIDS) []string {
	path := []string{}
	for node != nil {
		path = append(path, node.link)
		node = node.parent
	}
	return reverse(path)
}
