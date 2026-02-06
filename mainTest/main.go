package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("Creating.....")
	collector := colly.NewCollector(
		colly.AllowedDomains("127.0.0.1"),
	)

	collector.OnHTML(".category-item", func(h *colly.HTMLElement) {
		fmt.Println("tag name: ", h.Name, " index: ", h.Index, " text: ", h.Text)
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	collector.Visit("http://127.0.0.1:5500/mainTest/index.html")
}
