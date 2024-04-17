package logic

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

func BFS(linkMulai string, linkTujuan string, bahasa string) [][]string {
	queue := make(chan *Node, int(math.Pow(2, 24)))
	titleMulai := getPageTitle(linkMulai)
	start := newNode(linkMulai, titleMulai)
	hasil := [][]string{}
	titleVisited := make(map[string]bool)
	listHasil := []*Node{}
	found := false

	if bahasa == "ID" {
		pathUtama = pathUtamaIndo
	} else if bahasa == "EN" {
		pathUtama = pathUtamaInggris
	}

	titleTujuan := getPageTitle(linkTujuan)

	queue <- start
	titleVisited[titleMulai] = true

	if titleMulai == titleTujuan {
		return append(hasil, []string{linkMulai})
	}

	var wg sync.WaitGroup
	for len(queue) > 0 && !found {
		banyakNode := len(queue)
		node := <-queue

		aTags := getAllATag(node.link) // berisi map link dan title

		for _, aTag := range aTags {
			link := aTag["link"]
			title := aTag["title"]

			if titleVisited[title] {
				continue
			}

			if title == titleTujuan {
				nodeAkhir := newNode(link, title)
				nodeAkhir.parent = node
				listHasil = append(listHasil, nodeAkhir)
				found = true
			}

			if !titleVisited[title] {
				titleVisited[title] = true
				if !found {
					wg.Add(1)
					go func(n *Node) {
						defer wg.Done()
						n.parent = node
						queue <- n
					}(newNode(link, title))
					time.Sleep(time.Duration(rand.Intn(170)) * time.Microsecond)
				}
			}
		}
		time.Sleep(time.Duration(300) * time.Microsecond * time.Duration(banyakNode))
		wg.Wait()
	}

	for i := len(listHasil) - 1; i >= 0; i-- {
		hasil = append(hasil, getPath(listHasil[i]))
	}

	return hasil
}
