package data

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
)

type Response struct {

	Request *Request

	HttpHeader *http.Header
	HttpCookie *http.Cookie
	HttpResponse *http.Response

	StrBody string

	HtmlBody *goquery.Document
	JsonBody *simplejson.Json
}

