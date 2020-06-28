package main

import (
	"project/crawl/engine"
	"project/crawl/parse"
	"project/crawl/scheduler"
)

// go 爬虫项目
const domain = "https://book.douban.com"
func main() {
	req := scheduler.Request{
		Url: domain,
		ParseFunc: parse.GetTags,
	}
	//engine.Run(domain, req)

	e := engine.CoroutinesEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerNum: 100,
		Domain: domain,
	}

	e.Run(req)
}
