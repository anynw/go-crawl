package engine

import (
	"fmt"
	"github.com/anynw/go-crawl/fetch"
	"log"
)

type Scheduler interface {
	Submit(Request)
	configureWorkChan(chan Request)
}

type ConCurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type SimpleSchedule struct {
	workerChan chan Request
}

func (s *SimpleSchedule) Submit(request Request) {
	go func() { s.workerChan <- request }()
}

func (s *SimpleSchedule) configureWorkChan(requests chan Request) {
	s.workerChan = requests
}

func (e *ConCurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.configureWorkChan(in)

	for i := 0; i < e.WorkCount; i++ {
		CreateWork(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 处理out
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item :%d,%v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

func CreateWork(in chan Request, out chan ParseResult) {
	// 创建一个协程
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				fmt.Println("worker err :", err)
				continue
			}
			out <- result
		}
	}()
}

func worker(request Request) (ParseResult, error) {
	fmt.Printf("Fetch url:%s", request.Url)
	body, err := fetch.Fetch(request.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return request.ParseFunc(body), nil
}
