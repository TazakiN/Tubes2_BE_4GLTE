package logic

func BFS(linkMulai string, linkTujuan string) []string {
	queue := []*Node{}
	start := newNode(linkMulai)
	hasil := []string{}

	titleTujuan := getPageTitle(linkTujuan)

	queue = append(queue, start)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if getPageTitle(current.link) == titleTujuan {
			for current != nil {
				hasil = append(hasil, getPageTitle(current.link))
				current = current.parent
			}
			break
		}

		neighbours := getAllATag(current.link)
		for _, neighbourLink := range neighbours {
			neighbour := newNode(neighbourLink)
			if !neighbour.visited {
				neighbour.visited = true
				neighbour.distance = current.distance + 1
				neighbour.parent = current
				queue = append(queue, neighbour)
			}
		}
	}
	return hasil
}
