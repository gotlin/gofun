package data

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
)

type Response struct {

	request *Request

	Header *http.Header
	Cookie *http.Cookie

	StrBody string

	HtmlBody *goquery.Document
	JsonBody *simplejson.Json
}

