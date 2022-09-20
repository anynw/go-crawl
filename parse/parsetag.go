package parse

import (
	"regexp"

	"github.com/anynw/go-crawl/engine"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`

func ParseTag(content []byte) engine.ParseResult {
	re := regexp.MustCompile(regexpStr)
	matchs := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: "https://book.douban.com" + string(m[1]),
			// ParseFunc: engine.NilParse,
			ParseFunc: ParseBookList,
		})
	}

	return result

}
