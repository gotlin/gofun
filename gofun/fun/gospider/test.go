package main

import (
	"github.com/hu17889/go_spider/core/spider"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/PuerkitoBio/goquery"
)

type MyProcesser struct{}

func (mp *MyProcesser) Process(p *page.Page) {

	doc:=p.GetHtmlParser()
	doc.Find("a").Each(func(i int, chi *goquery.Selection) {

		href, _ := chi.Attr("href")

		p.AddTargetRequest(href,"html")
	})
}

func (mp *MyProcesser) Finish() {


}

func main() {

	p := MyProcesser{}

	spider.NewSpider(&p, "res").
		AddUrl("https://sale.fenqile.com/VlVdQklWVFVFSFxcWUFI/index.html", "html"). // start url, html is the responce type ("html" or "json" or "jsonp" or "text")
		AddPipeline(pipeline.NewPipelineConsole()). // Print result to std output
		SetSleepTime("rand", 100, 300). // Sleep time between 1s and 3s.
		Run()
}




