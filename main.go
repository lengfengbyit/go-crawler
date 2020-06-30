package main

import (
	"os"
	"project/crawl/engine"
	"project/crawl/parse/douban/book"
	book2 "project/crawl/persist/douban/book"
	"project/crawl/scheduler"
	"strconv"
)

// go 爬虫项目
const domain = "https://book.douban.com"
func main() {
	req := scheduler.Request{
		Url:       domain,
		ParseFunc: book.GetTags,
	}

	workerNum, err :=  strconv.Atoi(os.Getenv("WORKER_NUM"))
	if err != nil {
		panic(err)
	}

	e := engine.QueueEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerNum: workerNum,
		Domain:    domain,
		ItemChan:  book2.ItemSave(), // 爬取到的信息，持久化
	}

	e.Run(req)
}
