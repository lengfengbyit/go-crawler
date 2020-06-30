package book

import (
	"project/crawl/scheduler"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`
// 解析HTML, 获取标签
func GetTags(content []byte) scheduler.ParseResult {
	// <a href="/tag/漫画" class="tag">漫画</a>
	re := regexp.MustCompile(regexpStr)
	match := re.FindAllSubmatch(content, -1)

	results := scheduler.ParseResult{}
	for _, m := range match {
		results.Items = append(results.Items, m[2])
		results.Requests = append(results.Requests, scheduler.Request{
			Url:       string(m[1]),
			ParseFunc: GetBookList,
		})
	}
	return results
}
