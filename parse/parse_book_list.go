package parse

import (
	"project/crawl/scheduler"
	"regexp"
)

const bookListRegexpStr = `<a href="([^"]+)" title="([^"]+)"`

// 获取书列表
func GetBookList(content []byte) scheduler.ParseResult {

	re := regexp.MustCompile(bookListRegexpStr)
	match := re.FindAllSubmatch(content, -1)

	results := scheduler.ParseResult{}
	for _, m := range match {
		results.Items = append(results.Items, m[2])
		results.Requests = append(results.Requests, scheduler.Request{
			Url:       string(m[1]),
			ParseFunc: GetBookDetail,
		})
	}
	return results
}
