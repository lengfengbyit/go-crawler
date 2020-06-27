package parse

import (
	"project/crawl/engine"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`
// 解析HTML, 获取标签
func GetTags(content []byte) engine.ParseResult {
	// <a href="/tag/漫画" class="tag">漫画</a>
	re := regexp.MustCompile(regexpStr)
	match := re.FindAllSubmatch(content, -1)

	results := engine.ParseResult{}
	for _, m := range match {
		results.Items = append(results.Items, m[2])
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: nil,
		})
	}
	return results
}
