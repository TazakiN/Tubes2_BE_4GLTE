package logic

import (
	//"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var pathUtama string

const pathUtamaIndo = "https://id.wikipedia.org/wiki/"

const pathUtamaInggris = "https://en.wikipedia.org/wiki/"

func getPageTitle(url string) string {
	c := colly.NewCollector()

	title := ""

	c.OnHTML("span.mw-page-title-main", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.Visit(pathUtama + url)

	return title
}

func getAllATag(url string) []map[string]string {
	c := colly.NewCollector()

	links := []map[string]string{}

	c.OnHTML("div.mw-body-content a[href*='/wiki/'][title]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		title := e.Attr("title")

		if !strings.HasPrefix(href, "/wiki/") {
			return // skip yang gapunya /wiki/
		}

		if strings.Contains(href, "Berkas:") || strings.Contains(title, "Templat:") || strings.Contains(href, "Istimewa:") || strings.Contains(href, "Portal:") || strings.Contains(href, "Bantuan:") || strings.Contains(href, "Kategori:") || strings.Contains(href, "Wikipedia:") {
			return // skip berkas, templat, istimewa, portal, bantuan, kategori, wikipedia:
		}

		if strings.Contains(href, "File:") || strings.Contains(title, "Template:") || strings.Contains(href, "Special:") || strings.Contains(href, "Portal:") || strings.Contains(href, "Template_talk:") || strings.Contains(href, "Category:") {
			return // skip berkas, templat, istimewa, portal, bantuan, kategori, wikipedia:
		}

		if title == "" {
			return // skip yang gapunya title
		}

		if strings.HasPrefix(href, "/wiki/") {
			link := map[string]string{
				"link":  strings.TrimPrefix(href, "/wiki/"),
				"title": title,
			}
			links = append(links, link)
		}
	})

	c.Visit(pathUtama + url)

	return links
}
