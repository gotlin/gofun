package downloader

import (
	"gofun/fun/scrago/data"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
)

type HttpDownloader struct {
	client *http.Client
}

func (h *HttpDownloader) Download(req *data.Request) (*data.Response, error) {

	log.Println("downloading " + req.Url)

	httpReq, err := http.NewRequest("GET", req.Url, nil)
	if err != nil {
		log.Panicln(err)
	}

	httpResp, err := h.client.Do(httpReq)

	if err != nil {
		log.Panicln(err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, nil
	}

	doc, err := goquery.NewDocumentFromReader(httpResp.Body)

	if err != nil {
		log.Println(err)
	}

	defer httpResp.Body.Close()

	resp := &data.Response{}

	resp.StrBody = doc.Text()

	//log.Println("download text res " + resp.StrBody)

	resp.HtmlBody = doc

	resp.HttpResponse = httpResp

	resp.Request = req

	return resp, nil
}

func New() *HttpDownloader {

	return &HttpDownloader{&http.Client{}}

}
