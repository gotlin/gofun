package scheduler

import "gofun/fun/scrago/data"

type Scheduler interface{

	Poll() *data.Request
	Push(req *data.Request)

}