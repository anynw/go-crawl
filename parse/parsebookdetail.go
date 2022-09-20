package parse

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/anynw/go-crawl/engine"
	"github.com/anynw/go-crawl/model"
)

var autoRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a .*?>([^<]+)</a>`)
var publicRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\d\D]*?<a .*?>([^<]+)</a>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span>([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var intoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(contents []byte, bookName string) engine.ParseResult {
	// fmt.Printf("%s", contents)
	bookdetail := model.BookDetail{}

	bookdetail.BookName = bookName
	bookdetail.Author = ExtraString(contents, autoRe)
	bookdetail.Publicer = ExtraString(contents, publicRe)
	//类型转换
	str := ExtraString(contents, pageRe)
	// 去空格 strings.TrimSpace
	page, err := strconv.Atoi(strings.TrimSpace(str))
	if err == nil {
		bookdetail.Bookpages = page
	}

	bookdetail.Price = ExtraString(contents, priceRe)
	bookdetail.Score = ExtraString(contents, scoreRe)
	bookdetail.Into = ExtraString(contents, intoRe)

	result := engine.ParseResult{
		Items: []interface{}{bookdetail},
	}
	return result
}

func ExtraString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
