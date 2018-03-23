package processor

import "gofun/fun/scrago/data"

type Processor interface {
	Process(req *data.Request, resp *data.Response)
}
