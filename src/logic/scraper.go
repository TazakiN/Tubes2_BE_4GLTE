package logic

import (
	"strings"

	"github.com/gocolly/colly"
)

var pathUtama string

const pathUtamaIndo = "https://id.wikipedia.org/wiki/"

const pathUtamaInggris = "https://en.wikipedia.org/wiki/"

func getAllATag(url string) []string {
	// fmt.Println("Visiting", url)

	c := colly.NewCollector()

	links := []string{}

	c.OnHTML("div.mw-body-content p a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasPrefix(href, "/wiki/") {
			// Cetak judul halaman Wikipedia
			// fmt.Println(e.Text, "->", pathUtamaIndo+href)
			links = append(links, strings.TrimPrefix(href, "/wiki/"))
		}
	})

	c.Visit(pathUtama + url)

	return links
}

func getPageTitle(url string) string {
	c := colly.NewCollector()

	title := ""

	c.OnHTML("span.mw-page-title-main", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.Visit(pathUtama + url)
	// fmt.Println("sedang mengunjungi", pathUtama+url, "dengan judul", title)

	return title
}
