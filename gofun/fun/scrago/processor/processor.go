package processor

import (
	"gofun/fun/scrago/data"
	"gofun/fun/scrago/scheduler"
)

type Processor interface {
	Process(resp *data.Response, sch scheduler.Scheduler)
}


