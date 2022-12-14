package main

import (
	"github.com/anynw/go-crawl/engine"
	"github.com/anynw/go-crawl/parse"
	"github.com/anynw/go-crawl/scheduler"
)

// https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4
// https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C
func main() {

	eg := engine.ConCurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
	}
	eg.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})

	// eg := engine.ConCurrentEngine{
	// 	Scheduler: &engine.SimpleSchedule{},
	// 	WorkCount: 100,
	// }
	// eg.Run(engine.Request{
	// 	Url:       "https://book.douban.com",
	// 	ParseFunc: parse.ParseTag,
	// })

	// engine.Run(engine.Request{
	// 	//测试解析标签
	// 	// Url:       "https://book.douban.com",
	// 	// ParseFunc: parse.ParseTag,
	// 	//测试解析标题
	// 	// Url:       "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
	// 	// ParseFunc: parse.ParseBookList,
	//
	// 	//测试解析书籍详情页
	// 	// Url:       "https://book.douban.com/subject/30293801/",
	// 	// ParseFunc: parse.ParseBookDetail,
	//
	// 	//汇总版本，tag->booklist->bookdetail
	// 	Url:       "https://book.douban.com",
	// 	ParseFunc: parse.ParseTag,
	// })

	// 测试正则有效性
	// const str = `<div id="info" class="">`
	// re := regexp.MustCompile(str)
	// match := re.FindString(str)
	// fmt.Println("match:", match)

}
