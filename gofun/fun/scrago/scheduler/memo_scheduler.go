package scheduler

import "gofun/fun/scrago/data"

type Queue struct {
	Item *data.Request
	Next *Queue
}

type MemoScheduler struct {
	queue  *Queue
	fliter map[string]struct{}
}

func (m *MemoScheduler) Poll() *data.Request {

	if m.queue.Next != nil {
		req := m.queue.Next.Item
		m.queue.Next = m.queue.Next.Next
		return req
	}

	return nil

}

func (m *MemoScheduler) Push(req *data.Request) {

	if _, ok := m.fliter[req.Url]; ok {

		return
	}

	m.queue.Next = &Queue{req, nil}

	m.fliter[req.Url] = struct{}{}

}

func New() *MemoScheduler {

	return &MemoScheduler{&Queue{nil, nil}, make(map[string]struct{})}
}
