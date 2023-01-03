package parser

import (
	"fmt"

	"github.com/diyliv/cve/internal/models"
	"github.com/gocolly/colly/v2"
)

func ParseCveMitre(search string) []models.CVE {
	url := fmt.Sprintf("https://cve.mitre.org/cgi-bin/cvekey.cgi?keyword=%s", search)

	c := colly.NewCollector(colly.AllowURLRevisit())

	var data models.CVE
	var result []models.CVE

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			data.Name = h.ChildText("a")
			data.Link = "https://cve.mitre.org" + h.ChildAttr("a", "href")
			data.Description = h.ChildText("td")

			result = append(result, data)
		})
	})

	c.Visit(url)

	return result
}
