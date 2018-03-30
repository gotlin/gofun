package processor

import (
	"gofun/fun/scrago/data"
	"gofun/fun/scrago/scheduler"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"log"
)

type SimpleProcessor struct {
}

func (s *SimpleProcessor) Process(resp *data.Response, sch scheduler.Scheduler) {

	//html, _ := resp.HtmlBody.Html()
	//log.Println(html)

	resp.HtmlBody.Find("div a").Each(func(i int, chi *goquery.Selection) {

		href, _ := chi.Attr("href")

		log.Println("find href:" + href)

		link, _ := resp.HttpResponse.Request.URL.Parse(href)

		if strings.Contains(link.String(), "www.zhihu.com") {

			log.Println("parent url:" + resp.Request.Url + " add new:" + link.String())
			sch.Push(&data.Request{link.String()})

		}
	})

}

func New() *SimpleProcessor {
	return &SimpleProcessor{}
}
