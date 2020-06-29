package engine

// 协程版队列调度器爬虫

import (
	"log"
	"project/crawl/fetcher"
	"project/crawl/scheduler"
)

type QueueEngine struct {
	WorkerNum int                 // 协程数
	Scheduler scheduler.Scheduler // 调度器
	Domain    string              // 域名配置
	ItemChan  chan interface{}    // 需要保存的信息
}

func (engine *QueueEngine) Run(seeks ...scheduler.Request) {
	outChan := make(chan scheduler.ParseResult)

	// 启动调度器
	engine.Scheduler.Run()

	for i := 0; i < engine.WorkerNum; i++ {
		engine.CreateWorker(outChan)
	}

	// 初始化种子任务
	for _, r := range seeks {
		engine.Scheduler.Submit(r)
	}

	// 打印结果
	itemCount := 0
	for {
		itemCount++
		result := <-outChan
		for _, item := range result.Items {
			// 保存信息, 参数传递是为了防止变量引用传递的问题
			go func(tmp interface{}) {engine.ItemChan <- tmp}(item)
		}

		for _, req := range result.Requests {
			if req.Url[0:4] != "http" {
				req.Url = engine.Domain + req.Url
			}
			engine.Scheduler.Submit(req)
		}
	}
}

// 创建worker协程
func (engine *QueueEngine) CreateWorker(outChan chan scheduler.ParseResult) {

	go func() {
		// 从调度器中获取channel
		inChan := engine.Scheduler.GetWorkerChan()
		for {
			// 通知调度器，当前worker空闲，可以分派任务
			engine.Scheduler.WorkerReady(inChan)
			result, err := engine.Worker(<-inChan)
			if err != nil {
				log.Println(err)
				continue
			}
			outChan <- result
		}
	}()
}

func (engine *QueueEngine) Worker(r scheduler.Request) (result scheduler.ParseResult, err error) {

	log.Printf("Fetching Url: %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch Error: %v", err)
		return scheduler.ParseResult{}, err
	}

	return r.ParseFunc(body), nil
}
