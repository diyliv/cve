package parser

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func ParseCveMitre(search string) {
	url := fmt.Sprintf("https://cve.mitre.org/cgi-bin/cvekey.cgi?keyword=%s", search)

	c := colly.NewCollector(colly.AllowURLRevisit())

	c.OnHTML("table", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(url)
}
