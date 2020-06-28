package engine

import (
	"log"
	"project/crawl/fetcher"
	"project/crawl/scheduler"
)

// 运行爬虫程序
func Run(domain string, seeks ...scheduler.Request)  {

	var requests []scheduler.Request
	for _, e := range seeks {
		requests = append(requests, e)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		if r.Url[0:4] != "http" {
			r.Url = domain + r.Url
		}

		log.Printf("Fetching Url: %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch Error: %v", err)
			continue
		}

		if r.ParseFunc == nil {
			continue
		}

		result := r.ParseFunc(body)
		if len(result.Requests) > 0 {
			requests = append(requests, result.Requests...)
		}

		for _, item := range result.Items{
			log.Printf("Fetch item: %s", item)
		}
	}
}
