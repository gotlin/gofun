package scrago

import (
	"gofun/fun/scrago/scheduler"
	"gofun/fun/scrago/downloader"
	"gofun/fun/scrago/pipeline"
	"gofun/fun/scrago/processor"
)

type Scrago struct {
	sch  scheduler.Scheduler
	dow  downloader.Downloader
	pip  pipeline.Pipeline
	prc  processor.Processor
	pool chan struct{}
	dura int32
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

func (s *Scrago) Pool(curr int) *Scrago {
	//TODO 不能指针类型？
	s.pool = make(chan struct{}, curr)

	return s

}



func (s *Scrago) SetDefault() {
	if s.pool == nil {
		s.pool = make(chan struct{}, 10)
	}

	if s.dow == nil {
		s.Downloader(downloader.New())
	}

	if s.sch == nil{
		s.Scheduler(scheduler.New())
	}

	if s.pip == nil {
		s.Pipeline(pipeline.New())
	}




}
