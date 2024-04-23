package logic

import (
	//"encoding/json"
	"fmt"
	"math"
)

func IDS(linkMulai string, linkTujuan string, bahasa string, maxDepth int) ([][]string, [][]string) {
	// Create a slice to store the result
	hasil := [][]string{}
	visit := [][]string{}

	// Loop through increasing depth limits
	for depth := 0; depth <= maxDepth; depth++ {
		// Perform depth-limited search with current depth limit
		hasil, visit = DLS(linkMulai, linkTujuan, bahasa, depth)

		// If the destination is found, return the result
		if len(hasil) > 0 {
			return hasil, visit
		}
	}

	// Return empty result if destination not found within max depth
	return hasil, visit
}

func DLS(linkMulai string, linkTujuan string, bahasa string, depth int) ([][]string, [][]string) {
	// Create a stack for the queue
	queue := make(chan *NodeIDS, int(math.Pow(2, 24)))

	// Get the title of the starting page
	titleMulai := getPageTitle(linkMulai)

	// Create the starting node
	start := newNodeIDS(linkMulai, titleMulai, 0) // Add depth parameter to newNode

	// Create a slice to store the result
	hasil := [][]string{}
	visit := [][]string{}

	// Create a map to keep track of visited titles
	titleVisited := NewSafeTitleVisited()

	// Create a slice to store the list of result nodes
	listHasil := []*NodeIDS{}

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

	// Process stack until empty or depth limit is reached
	for len(queue) > 0 {
		// Dequeue a node from the stack
		node := <-queue
		fmt.Println("Title visited: " + node.title)

		// Check if depth limit is reached
		if node.depth == depth {
			continue // Skip expanding this node further
		}

		// Get all <a> tags from the node's link
		aTags := getAllATag(node.link)

		// Iterate over each <a> tag
		for _, aTag := range aTags {
			link := aTag["link"]
			title := aTag["title"]

			// Append the title to the visit list
			visit = append(visit, []string{title})

			// Check if the title has been visited before
			if titleVisited.HasVisited(title) {
				continue
			}

			// Check if the destination title is found
			if title == titleTujuan {
				// Create a new node for the destination
				nodeAkhir := newNodeIDS(link, title, node.depth+1) // Increment depth
				nodeAkhir.parent = node

				// Append the destination node to the list of result nodes
				listHasil = append(listHasil, nodeAkhir)

				// Return the result
				for i := len(listHasil) - 1; i >= 0; i-- {
					hasil = append(hasil, getPathIDS(listHasil[i]))
				}
				return hasil, visit
			}

			// Mark the title as visited
			titleVisited.MarkVisited(title)

			// Enqueue the new node
			n := newNodeIDS(link, title, node.depth+1) // Increment depth
			n.parent = node
			queue <- n
		}
	}

	// Return empty result if destination not found within depth limit
	return hasil, visit
}

// Function signature for DLS:
// func DLS(linkMulai string, linkTujuan string, bahasa string, depth int) ([][]string, [][]string) {
//     // Implement Depth-Limited Search (DLS) here
// }

// type Result struct {
// 	Method     string `json:"method"`
// 	LinkAwal   string `json:"linkAwal"`
// 	LinkTujuan string `json:"linkTujuan"`
// }

// func IDS(linkMulai string, linkTujuan string, bahasa string) [][]string {
// 	result := Result{
// 		Method:     "IDS",
// 		LinkAwal:   linkMulai,
// 		LinkTujuan: linkTujuan,
// 	}

// 	// proses pencarian jalur di sini

// 	jsonResult, err := json.Marshal(result)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return nil
// 	}

// 	fmt.Println(string(jsonResult))

// 	return nil
// }
