package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Item struct {
	Name string
	Desc string
	Url  string
}

func (item *Item) Print() {
	fmt.Println(item.Name, item.Url)
	fmt.Println(item.Desc)
	fmt.Println("=============================")
	fmt.Println()
}

func search(q string) []*Item {

	var result []*Item
	c := colly.NewCollector()

	replacer := strings.NewReplacer("(", "", ")", "")

	c.OnHTML(".SearchSnippet", func(e *colly.HTMLElement) {
		text := e.ChildText("h2 > a")

		a := strings.ReplaceAll(text, " ", "")
		sp := strings.Split(a, "\n")
		name := sp[0]
		url := replacer.Replace(sp[1])

		desc := e.ChildText(".SearchSnippet-synopsis")

		item := &Item{
			Name: name,
			Url:  url,
			Desc: desc,
		}

		result = append(result, item)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	url := fmt.Sprintf("https://pkg.go.dev/search?q=%s&m=package&limit=5", q)
	c.Visit(url)

	return result
}
