package model

import "strconv"

type BookDetail struct {
	Author    string
	Publicer  string
	Bookpages int
	Price     string
	// 书籍评分
	Score string
	// 内容简介
	Into string
}

func (b BookDetail) String() string {
	return "作者：" + b.Author + "\n出版社：" + b.Publicer + "\n页数：" + strconv.Itoa(b.Bookpages) + "\n价格：" + b.Price + "\n得分：" + b.Score + "\n内容简介：" + b.Into
}
