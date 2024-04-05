package logic

func BFS(linkMulai string, linkTujuan string, bahasa string) []string {
	queue := []*Node{}
	titleMulai := getPageTitle(linkMulai)
	start := newNode(linkMulai, titleMulai)
	hasil := []string{}
	titleVisited := make(map[string]bool)
	node := &Node{}
	found := false

	if bahasa == "id" {
		pathUtama = pathUtamaIndo
	} else if bahasa == "en" {
		pathUtama = pathUtamaInggris
	}

	titleTujuan := getPageTitle(linkTujuan)

	queue = append(queue, start)
	titleVisited[titleMulai] = true

	if titleMulai == titleTujuan {
		return append(hasil, titleMulai)
	}

	for len(queue) > 0 && !found {
		current := queue[0]
		// fmt.Println("visiting", current.title, "di link", current.link)
		queue = queue[1:]

		aTags := getAllATag(current.link) // berisi map link dan title

		for _, aTag := range aTags {
			link := aTag["link"]
			title := aTag["title"]

			if titleVisited[title] {
				continue
			}

			if title == titleTujuan {
				node = newNode(link, title)
				node.parent = current

				found = true
				break
			}

			if !titleVisited[title] {
				titleVisited[title] = true
				neighbour := newNode(link, title)
				neighbour.parent = current
				current.neighbours = append(current.neighbours, neighbour)
				queue = append(queue, neighbour)
			}
		}
	}

	if found {
		for node != nil {
			hasil = append(hasil, node.title)
			node = node.parent
		}
		hasil = reverse(hasil)
	}
	return hasil
}
