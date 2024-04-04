package logic

import (
	"strings"

	"github.com/gocolly/colly"
)

const pathUtamaIndo = "https://id.wikipedia.org/wiki/"

// const pathUtamaInggris = "https://en.wikipedia.org"

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

	c.Visit(pathUtamaIndo + url)

	return links
}

func getPageTitle(url string) string {
	c := colly.NewCollector()

	title := ""

	c.OnHTML("span.mw-page-title-main", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.Visit(pathUtamaIndo + url)

	return title
}
