package logic

import (
	//"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
	"time"
)

func IDS(linkMulai string, linkTujuan string, bahasa string, maxDepth int) ([][]string, [][]string) {
	// Create a slice to store the result
	hasil := [][]string{}
	visit := [][]string{}

	//Check starting page
	hasil, visit = DLSLevelZero(linkMulai, linkTujuan, bahasa)

	// If the destination is found, return the result
	if len(hasil) > 0 {
		return hasil, visit
	}

	// Loop through increasing depth limits
	for depth := 1; depth <= maxDepth; depth++ {
		// Perform depth-limited search with current depth limit
		hasil, visit = DLS(linkMulai, linkTujuan, bahasa, depth)

		// If the destination is found, return the result
		if len(hasil) > 0 {
			return hasil, visit
		}
	}

	// Return empty result if destination not found within max depth
	fmt.Println("Not found! Try searching deeper")
	return hasil, visit
}

func DLSLevelZero(linkMulai string, linkTujuan string, bahasa string) ([][]string, [][]string) {
	// Get the title of the starting page
	titleMulai := getPageTitle(linkMulai)

	start := newNodeIDS(linkMulai, titleMulai, 0)

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

	titleVisited.MarkVisited(titleMulai)

	// If the starting and destination titles are the same, return the link
	if titleMulai == titleTujuan {
		return append(hasil, []string{linkMulai}), visit
	}

	aTags := getAllATag(start.link)

	// Iterate over each <a> tag
	for _, aTag := range aTags {
		link := aTag["link"]
		title := aTag["title"]

		// Check if the title has been visited before
		if titleVisited.HasVisited(title) {
			continue
		}

		visit = append(visit, []string{title})

		// Check if the destination title is found
		if title == titleTujuan {
			nodeAkhir := newNodeIDS(link, title, start.depth+1)
			nodeAkhir.parent = start
			listHasil = append(listHasil, nodeAkhir)
			break
		}

		// Mark the title as visited
		titleVisited.MarkVisited(title)
	}

	// Generate the path from the result nodes
	for i := len(listHasil) - 1; i >= 0; i-- {
		hasil = append(hasil, getPathIDS(listHasil[i]))
	}

	// Return the result
	return hasil, visit
}

func DLS(linkMulai string, linkTujuan string, bahasa string, depth int) ([][]string, [][]string) {
	// Create channels for results
	//resultChan := make(chan [][]string)
	//visitChan := make(chan [][]string)

	// Create a stack for the queue
	queue := make(chan *NodeIDS, int(math.Pow(2, 24)))

	// Get the title of the starting page
	titleMulai := getPageTitle(linkMulai)

	// Create the starting node
	start := newNodeIDS(linkMulai, titleMulai, 0)

	// Create a slice to store the result
	hasil := [][]string{}
	visit := [][]string{}

	// Create a map to keep track of visited titles
	titleVisited := NewSafeTitleVisited()

	// Determine the path based on the language
	if bahasa == "ID" {
		pathUtama = pathUtamaIndo
	} else if bahasa == "EN" {
		pathUtama = pathUtamaInggris
	}

	// Get the title of the destination page
	titleTujuan := getPageTitle(linkTujuan)

	// Queue all <a> tags from starting page
	aTags := getAllATag(start.link)

	// Iterate over each <a> tag
	for _, aTag := range aTags {
		link := aTag["link"]
		title := aTag["title"]

		// Check if the title has been visited before
		if titleVisited.HasVisited(title) {
			continue
		}

		// Mark the title as visited
		titleVisited.MarkVisited(title)

		n := newNodeIDS(link, title, start.depth+1)
		n.parent = start
		queue <- n
	}

	found := false

	var wg sync.WaitGroup
	// Process stack until empty or depth limit is reached
	for len(queue) > 0 {
		// Dequeue a node from the stack
		node := <-queue
		fmt.Println("Title visited: " + node.title)
		fmt.Printf("Length of queue now: %d\n", len(queue))
		fmt.Printf("Found? %t\n", found)
		fmt.Printf("Depth %d\n", depth)

		// Check if depth limit is reached
		if node.depth > depth {
			fmt.Println("error sini")
			continue // Skip expanding this node further
		}

		time.Sleep(time.Duration(rand.IntN(21)) * time.Millisecond)

		wg.Add(1)
		go func(nodeToProcess *NodeIDS) {
			// Decrement the WaitGroup counter when the goroutine finishes
			defer wg.Done()

			//var localHasil [][]string

			// Get all <a> tags from the node's link
			aTags := getAllATag(nodeToProcess.link)

			// Iterate over each <a> tag
			for _, aTag := range aTags {
				link := aTag["link"]
				title := aTag["title"]

				if found {
					break
				}

				// Check if the title has been visited before
				if titleVisited.HasVisited(title) {
					continue
				}

				visit = append(visit, []string{node.title})

				// Check if the destination title is found
				if title == titleTujuan {
					// Create a new node for the destination
					nodeAkhir := newNodeIDS(link, title, nodeToProcess.depth+1) // Increment depth
					nodeAkhir.parent = nodeToProcess

					// Append the destination node to the list of result nodes
					// Append the destination node to the list of result nodes
					hasil = append(hasil, getPathIDS(nodeAkhir))

					found = true

					break
				}

				// Mark the title as visited
				titleVisited.MarkVisited(title)

				if !found {
					// Enqueue the new node
					n := newNodeIDS(link, title, nodeToProcess.depth+1) // Increment depth
					n.parent = nodeToProcess
					queue <- n
				}
			}
		}(node)
	}

	fmt.Println("stuck di wait")
	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("ga di wait")
	// Receive results from channels if found
	// if found {
	// 	for res := range resultChan {
	// 		hasil = append(hasil, res...)
	// 	}
	// }

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
