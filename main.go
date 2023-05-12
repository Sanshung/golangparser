package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	var url string
	fmt.Println("Введите урл книги с сайта rosman.ru")
	fmt.Scanf("%s\n", &url)

	c := colly.NewCollector()
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnHTML(".detailtext", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnHTML("#item-pic", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("src"))
	})
	c.OnHTML(".ozon", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})
	c.OnHTML("#chars", func(e *colly.HTMLElement) {
		e.ForEach("dt", func(_ int, el *colly.HTMLElement) {
			e.ForEach("dd", func(_ int, el2 *colly.HTMLElement) {
				if el.Index == el2.Index {
					fmt.Println(el.Text + ": " + el2.Text)
				}
			})
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit(url)
}
