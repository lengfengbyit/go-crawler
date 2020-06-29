package main

import (
	"os"
	"project/crawl/engine"
	"project/crawl/parse"
	"project/crawl/persist"
	"project/crawl/scheduler"
	"strconv"
)

// go 爬虫项目
const domain = "https://book.douban.com"
func main() {
	req := scheduler.Request{
		Url: domain,
		ParseFunc: parse.GetTags,
	}
	//engine.Run(domain, req)

	//e := engine.CoroutinesEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkerNum: 10,
	//	Domain: domain,
	//}

	workerNum, err :=  strconv.Atoi(os.Getenv("WORKER_NUM"))
	if err != nil {
		panic(err)
	}

	e := engine.QueueEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerNum: workerNum,
		Domain: domain,
		ItemChan: persist.ItemSave(), // 爬取到的信息，持久化
	}

	e.Run(req)
}
