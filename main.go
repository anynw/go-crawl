package main

import (
	"github.com/anynw/go-crawl/engine"
	"github.com/anynw/go-crawl/parse"
)

//https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4
//https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C
func main() {
	engine.Run(engine.Request{
		//测试解析标签
		// Url:       "https://book.douban.com",
		// ParseFunc: parse.ParseTag,
		//测试解析标题
		Url:       "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
		ParseFunc: parse.ParseBookList,
	})
}
