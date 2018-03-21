package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://www.baidu.com/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(doc.Text())

	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		href, _ := s.Attr("href")

		fmt.Println(s.Text())
		fmt.Println("href:"+href)
	})
}

func main() {
	ExampleScrape()
}