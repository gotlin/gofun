package scrago

import (
	"gofun/fun/scrago/scheduler"
	"gofun/fun/scrago/downloader"
	"gofun/fun/scrago/pipeline"
	"gofun/fun/scrago/processor"
	"gofun/fun/scrago/goropool"
	"gofun/fun/scrago/data"
	"time"
	"log"
)

type Scrago struct {
	sch  scheduler.Scheduler
	dow  downloader.Downloader
	pip  pipeline.Pipeline
	prc  processor.Processor
	pool goropool.GoroPool
}

func New() *Scrago {
	return &Scrago{}
}

func (s *Scrago) Scheduler(sch scheduler.Scheduler) *Scrago {
	s.sch = sch
	return s
}

func (s *Scrago) Downloader(dow downloader.Downloader) *Scrago {
	s.dow = dow
	return s
}

func (s *Scrago) Pipeline(pip pipeline.Pipeline) *Scrago {
	s.pip = pip
	return s
}

func (s *Scrago) Processor(prc processor.Processor) *Scrago {

	s.prc = prc
	return s
}

func (s *Scrago) GoroPool(curr int) *Scrago {

	s.pool = goropool.New(curr)

	return s

}

func (s *Scrago) SetDefault() *Scrago {

	if s.pool == nil {
		s.GoroPool(10)
	}

	if s.dow == nil {
		s.Downloader(downloader.New())
	}

	if s.sch == nil {
		s.Scheduler(scheduler.New())
	}

	if s.pip == nil {
		s.Pipeline(pipeline.New())
	}

	if s.prc == nil {
		s.Processor(processor.New())
	}

	return s

}

func (s *Scrago) AddUrl(url string) *Scrago {

	s.sch.Push(&data.Request{url})

	return s
}

func (s *Scrago) Start() {

	for {

		time.Sleep(500 * time.Millisecond)

		request := s.sch.Poll()

		if request == nil {
			log.Println(".......none url")
			continue
		}

		f := func() {

			resp, _ := s.dow.Download(request)

			s.prc.Process(resp, s.sch)
		}

		s.pool.Submit(f)

	}

}
