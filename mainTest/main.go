package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("Creating.....")
	collector := colly.NewCollector(
		colly.AllowedDomains("guides.rubyonrails.org", "www.guides.rubyonrails.org"),
	)

	collector.OnHTML("header .wrapper ul li", func(h *colly.HTMLElement) {
		fmt.Println("tag name: ", h.Name, " index: ", h.Index, " text: ", h.Text)

	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	collector.Visit("https://guides.rubyonrails.org/active_record_validations.html")
}
