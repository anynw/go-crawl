package scheduler

import "github.com/anynw/go-crawl/engine"

type SimpleSchedule struct {
	workerChan chan engine.Request
}

func (s *SimpleSchedule) Submit(request engine.Request) {
	go func() { s.workerChan <- request }()
}

func (s *SimpleSchedule) configureWorkChan(requests chan engine.Request) {
	s.workerChan = requests
}
