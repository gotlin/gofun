package downloader

import (
	"gofun/fun/scrago/data"
	"net/http"
	"log"
	"io/ioutil"
)

type HttpDownloader struct {
	client *http.Client
}

func (h *HttpDownloader) Download(req *data.Request) (*data.Response, error) {

	httpReq, err := http.NewRequest("GET", req.Url, nil)
	if err != nil {
		log.Panicln(err)
	}

	httpResp, err := h.client.Do(httpReq)

	if err != nil {
		log.Panicln(err)
	}

	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		log.Panicln(err)
	}

	defer httpResp.Body.Close()

	resp := &data.Response{}

	resp.StrBody = string(body)

	return resp, nil
}

func New() *HttpDownloader {

	return &HttpDownloader{&http.Client{}}

}
