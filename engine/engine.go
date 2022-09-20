package engine

import (
	"fmt"
	"log"

	"github.com/anynw/go-crawl/fetch"
)

func Run(seeds ...Request) {
	var requests []Request

	// for _, e := range seeds {
	// 	requests = append(requests, e)
	// }

	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		// 踢除第一个元素
		requests = requests[1:]
		log.Printf("Fetching url : %s", r.Url)

		body, err := fetch.Fetch(r.Url)

		if err != nil {
			log.Printf("Fetch error : %s", r.Url)
		}

		pr := r.ParseFunc(body)
		requests = append(requests, pr.Requests...)

		for _, item := range pr.Items {
			fmt.Printf("Got item:%s\n", item)
		}

	}
}
