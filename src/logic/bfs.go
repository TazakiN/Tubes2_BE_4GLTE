package logic

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
	"time"
)

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

// BFS performs a breadth-first search to find paths between two Wikipedia pages.
func BFS(linkMulai string, linkTujuan string, bahasa string) ([][]string, [][]string) {
	// Create a channel for the queue
	queue := make(chan *Node, int(math.Pow(2, 24)))

	// Get the title of the starting page
	titleMulai := getPageTitle(linkMulai)

	// Create the starting node
	start := newNode(linkMulai, titleMulai)

	// Create a slice to store the result
	hasil := [][]string{}
	visit := [][]string{}

	// Create a map to keep track of visited titles
	titleVisited := NewSafeTitleVisited()

	// Create a slice to store the list of result nodes
	listHasil := []*Node{}

	// Variable to track if the destination is found
	found := false

	// Determine the path based on the language
	if bahasa == "ID" {
		pathUtama = pathUtamaIndo
	} else if bahasa == "EN" {
		pathUtama = pathUtamaInggris
	}

	// Get the title of the destination page
	titleTujuan := getPageTitle(linkTujuan)

	// Enqueue the starting node
	queue <- start

	titleVisited.MarkVisited(titleMulai)

	// If the starting and destination titles are the same, return the link
	if titleMulai == titleTujuan {
		return append(hasil, []string{linkMulai}), visit
	}

	//Process main page
	// Dequeue a node from the queue
	node := <-queue
	fmt.Println("Node visited: " + node.title)

	// Get all <a> tags from the node's link
	aTags := getAllATag(node.link)

	// Iterate over each <a> tag
	for _, aTag := range aTags {
		link := aTag["link"]
		title := aTag["title"]

		visit = append(visit, []string{title})

		// Check if the title has been visited before
		if titleVisited.HasVisited(title) {
			continue
		}

		// Check if the destination title is found
		if title == titleTujuan {
			nodeAkhir := newNode(link, title)
			nodeAkhir.parent = node
			listHasil = append(listHasil, nodeAkhir)
			found = true
			break
		}

		// Mark the title as visited
		titleVisited.MarkVisited(title)

		// If the destination is not found, enqueue the new node and start a goroutine
		if !found {
			n := newNode(link, title)
			n.parent = node
			queue <- n
		}
	}

	//Process queue if not found on main page
	//var mutex sync.Mutex
	var wg sync.WaitGroup
	// Loop until the queue is empty or the destination is found
	for len(queue) > 0 && !found {
		// Dequeue a node from the queue
		node := <-queue
		fmt.Println("Title visited: " + node.title)
		fmt.Printf("Length of queue now: %d\n", len(queue))
		fmt.Printf("Found? %t\n", found)

		time.Sleep(time.Duration(rand.IntN(21)) * time.Millisecond)
		// Increment the WaitGroup counter
		wg.Add(1)
		// Process each node concurrently using a goroutine
		go func(nodeToProcess *Node) {
			// Decrement the WaitGroup counter when the goroutine finishes
			defer wg.Done()
			// Get all <a> tags from the node's link
			aTags := getAllATag(nodeToProcess.link)

			// Iterate over each <a> tag
			for _, aTag := range aTags {
				link := aTag["link"]
				title := aTag["title"]

				if found {
					break
				}

				// Append the title to the visit list
				visit = append(visit, []string{title})

				// Check if the title has been visited before
				if titleVisited.HasVisited(title) {
					continue
				}

				// Check if the destination title is found
				if title == titleTujuan {
					// Create a new node for the destination
					nodeAkhir := newNode(link, title)
					nodeAkhir.parent = nodeToProcess

					// Append the destination node to the list of result nodes
					listHasil = append(listHasil, nodeAkhir)

					// Set found to true to break the loop
					found = true

					// Break the loop
					break
				}

				// Mark the title as visited
				titleVisited.MarkVisited(title)

				// If the destination is not found, enqueue the new node
				if !found {
					n := newNode(link, title)
					n.parent = nodeToProcess
					queue <- n
				}
			}
		}(node)
	}

	//close(queue)
	fmt.Printf("Found final? %t\n", found)

	// var closeOnce sync.Once
	// closeOnce.Do(func() {
	// 	close(queue)
	// })

	wg.Wait()

	// Generate the path from the result nodes
	for i := len(listHasil) - 1; i >= 0; i-- {
		hasil = append(hasil, getPath(listHasil[i]))
	}

	// Return the result
	return hasil, visit
}

//19:33 23/04/2024
//-stuck pas antara dia found true trus keluar loop tapi ada goroutine yang blm kelar
//-atau stuck gr2 len(queue) = 0 tapi blm found, possibly gara2 goroutinenya lambat jadi keduluan abis

//start title
//bikin path (array of string) = start title dulu
//array of path = simpen semua kemungkinan path -> udah discrape
//

// package logic

// import (
// 	//"fmt"
// 	"math"
// )

// // BFS performs a breadth-first search to find paths between two Wikipedia pages.
// func BFS(linkMulai string, linkTujuan string, bahasa string) ([][]string, [][]string) {
// 	// Create a channel for the queue
// 	queue := make(chan *Node, int(math.Pow(2, 24)))

// 	// Get the title of the starting page
// 	titleMulai := getPageTitle(linkMulai)

// 	// Create the starting node
// 	start := newNode(linkMulai, titleMulai)

// 	// Create a slice to store the result
// 	hasil := [][]string{}
// 	visit := [][]string{}

// 	// Create a map to keep track of visited titles
// 	titleVisited := make(map[string]bool)

// 	// Create a slice to store the list of result nodes
// 	listHasil := []*Node{}

// 	// Variable to track if the destination is found
// 	found := false

// 	// Determine the path based on the language
// 	if bahasa == "ID" {
// 		pathUtama = pathUtamaIndo
// 	} else if bahasa == "EN" {
// 		pathUtama = pathUtamaInggris
// 	}

// 	// Get the title of the destination page
// 	titleTujuan := getPageTitle(linkTujuan)

// 	// Enqueue the starting node
// 	queue <- start

// 	titleVisited[titleMulai] = true

// 	// If the starting and destination titles are the same, return the link
// 	if titleMulai == titleTujuan {
// 		return append(hasil, []string{linkMulai}), visit
// 	}

// 	// Loop until the queue is empty or the destination is found
// 	for len(queue) > 0 && !found {
// 		// Dequeue a node from the queue
// 		node := <-queue

// 		// Get all <a> tags from the node's link
// 		aTags := getAllATag(node.link)

// 		// Iterate over each <a> tag
// 		for _, aTag := range aTags {
// 			link := aTag["link"]
// 			title := aTag["title"]

// 			visit = append(visit, []string{title})

// 			visited := titleVisited[title]

// 			// Skip if the title has been visited before
// 			if visited {
// 				continue
// 			}

// 			// Check if the destination title is found
// 			if title == titleTujuan {
// 				nodeAkhir := newNode(link, title)
// 				nodeAkhir.parent = node
// 				listHasil = append(listHasil, nodeAkhir)
// 				found = true
// 				break
// 			}

// 			titleVisited[title] = true

// 			// If the destination is not found, enqueue the new node and start a goroutine
// 			if !found {
// 				n := newNode(link, title)
// 				n.parent = node
// 				queue <- n
// 			}
// 		}
// 	}

// 	// Generate the path from the result nodes
// 	for i := len(listHasil) - 1; i >= 0; i-- {
// 		hasil = append(hasil, getPath(listHasil[i]))
// 	}

// 	// Return the result
// 	return hasil, visit
// }
