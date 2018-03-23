package pipeline

import "gofun/fun/scrago/data"

type Pipeline interface{
	Pipe(resp *data.Response)
}
