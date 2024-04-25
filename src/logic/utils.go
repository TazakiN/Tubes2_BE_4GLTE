package logic

import (
	"fmt"
	"sync"
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

// SafeTitleVisited is a thread-safe map to keep track of visited titles
type SafeTitleVisited struct {
	mu     sync.Mutex
	titles map[string]bool
}

// NewSafeTitleVisited creates a new instance of SafeTitleVisited
func NewSafeTitleVisited() *SafeTitleVisited {
	return &SafeTitleVisited{
		titles: make(map[string]bool),
	}
}

// MarkVisited marks a title as visited
func (stv *SafeTitleVisited) MarkVisited(title string) {
	stv.mu.Lock()
	defer stv.mu.Unlock()
	stv.titles[title] = true
}

// HasVisited checks if a title has been visited
func (stv *SafeTitleVisited) HasVisited(title string) bool {
	stv.mu.Lock()
	defer stv.mu.Unlock()
	return stv.titles[title]
}
