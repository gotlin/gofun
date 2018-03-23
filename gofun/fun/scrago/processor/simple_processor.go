package processor

import (
	"gofun/fun/scrago/data"
	"gofun/fun/scrago/scheduler"
)

type SimpleProcessor struct {
}

func (s *SimpleProcessor) Process(resp *data.Response, sch scheduler.Scheduler) {

	body := resp.StrBody

	sch.Push(&data.Request{body})

}
