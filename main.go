package main

import (
	"project/crawl/engine"
	"project/crawl/parse"
)

// go 爬虫项目
const domain = "https://book.douban.com"
func main() {
	req := engine.Request{
		Url: "https://book.douban.com/subject/4913064/",
		ParseFunc: parse.GetBookDetail,
	}
	engine.Run(domain, req)
}
