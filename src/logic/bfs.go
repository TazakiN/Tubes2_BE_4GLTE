package logic

func BFS(linkMulai string, linkTujuan string, bahasa string) []string {
	queue := []*Node{}
	start := newNode(linkMulai)
	hasil := []string{}
	titleVisited := make(map[string]bool)

	if bahasa == "id" {
		pathUtama = pathUtamaIndo
	} else if bahasa == "en" {
		pathUtama = pathUtamaInggris
	}

	titleTujuan := getPageTitle(linkTujuan)

	queue = append(queue, start)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		titleCurrent := getPageTitle(current.link)

		if titleCurrent == titleTujuan {
			for current != nil {
				hasil = append(hasil, getPageTitle(current.link))
				current = current.parent
			}
			break
		}

		neighbours := getAllATag(current.link)
		for _, neighbourLink := range neighbours {
			if !titleVisited[neighbourLink] {
				neighbour := newNode(neighbourLink)
				neighbour.distance = current.distance + 1
				neighbour.parent = current
				queue = append(queue, neighbour)
				titleVisited[neighbourLink] = true
			}
		}
	}
	return hasil
}
