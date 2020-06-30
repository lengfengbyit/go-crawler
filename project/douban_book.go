package project

import (
	"os"
	"project/crawl/engine"
	"project/crawl/parse/douban/book"
	persist "project/crawl/persist/douban/book"
	"project/crawl/scheduler"
	"strconv"
)

func DoubanBook() {
	var domain = "https://book.douban.com"
	req := scheduler.Request{
		Url:       domain,
		ParseFunc: book.GetTags,
	}

	workerNum, err := strconv.Atoi(os.Getenv("WORKER_NUM"))
	if err != nil {
		panic(err)
	}

	e := engine.QueueEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerNum: workerNum,
		Domain:    domain,
		ItemChan:  persist.ItemSave(), // 爬取到的信息，持久化
	}
	e.Run(req)
}
