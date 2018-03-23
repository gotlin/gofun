package downloader

import "gofun/fun/scrago/data"

type Downloader interface {
	Download(req *data.Request) (*data.Response, error)
}
