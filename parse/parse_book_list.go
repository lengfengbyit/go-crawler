package parse

import (
	"project/crawl/engine"
	"regexp"
)

const bookListRegexpStr = `<a href="([^"]+)" title="([^"]+)"`

// 获取书列表
func GetBookList(content []byte) engine.ParseResult {

	re := regexp.MustCompile(bookListRegexpStr)
	match := re.FindAllSubmatch(content, -1)

	results := engine.ParseResult{}
	for _, m := range match {
		results.Items = append(results.Items, m[2])
		results.Requests = append(results.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: nil,
		})
	}
	return results
}
