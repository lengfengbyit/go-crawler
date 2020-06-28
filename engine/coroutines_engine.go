package engine

// 协程版爬虫

import (
	"log"
	"project/crawl/fetcher"
	"project/crawl/scheduler"
)

type CoroutinesEngine struct {
	WorkerNum int                 // 协程数
	Scheduler scheduler.Scheduler // 调度器
	Domain string // 域名配置
}

func (engine *CoroutinesEngine) Run(seeks ...scheduler.Request) {
	inChan := make(chan scheduler.Request)
	outChan := make(chan scheduler.ParseResult)

	engine.Scheduler.ConfigureWorkChan(inChan)

	for i := 0; i < engine.WorkerNum; i++ {
		engine.CreateWorker(inChan, outChan)
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
			log.Printf("%d, %s\n",itemCount, item)
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
func (engine *CoroutinesEngine) CreateWorker(inChan chan scheduler.Request, outChan chan scheduler.ParseResult) {

	go func() {
		for {
			result, err := engine.Worker(<-inChan)
			if err != nil {
				log.Println(err)
				continue
			}
			outChan <- result
		}
	}()
}

func (engine *CoroutinesEngine) Worker(r scheduler.Request) (result scheduler.ParseResult, err error) {

	log.Printf("Fetching Url: %s: %d\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch Error: %v", err)
		return scheduler.ParseResult{}, err
	}

	if r.ParseFunc == nil {
		return scheduler.ParseResult{Items: []interface{}{body}}, nil
	}
	return r.ParseFunc(body), nil
}
