package parse

import (
	"regexp"

	"github.com/anynw/go-crawl/engine"
)

//const BookListRe = `<a href="https://book.douban.com/subject/35680099/" title="偶像失格" onclick="moreurl(this,{i:'1',query:'',subject_id:'35680099',from:'book_subject_search'})">偶像失格</a>`
const BookListRe = `<a href="([^"]+)" title="([^"]+)"`

func ParseBookList(contents []byte) engine.ParseResult {
	// fmt.Printf("%s", contents)
	re := regexp.MustCompile(BookListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matchs {
		bookName := string(m[2])
		result.Items = append(result.Items, bookName)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// ParseFunc: ParseBookDetail,
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseBookDetail(c, bookName)
			},
		})
	}

	return result
}
