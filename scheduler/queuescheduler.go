package scheduler

import "github.com/anynw/go-crawl/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	// 二维通道
	workerChan chan chan engine.Request
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (s *QueueScheduler) WorkReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workQ = append(workQ, w)
			case activeWork <- activeRequest:
				// 删除第一个
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
