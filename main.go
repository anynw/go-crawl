package main

import (
	"github.com/anynw/go-crawl/engine"
	"github.com/anynw/go-crawl/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseContent,
	})
}
