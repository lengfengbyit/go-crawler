package scheduler

// 队列调度器，每个Worker都有一个队列
type QueueScheduler struct {
	RequestChan chan Request      // Request 类型的 chan
	WorkerChan  chan chan Request // Request chan 类型的 chan
}

func (q *QueueScheduler) Submit(request Request) {
	q.RequestChan <- request
}

func (q *QueueScheduler) ConfigureWorkChan(requests chan Request) {
	q.RequestChan = requests
}

// Worker提交chan Request到这里
func (q *QueueScheduler) WorkerReady(req chan Request) {
	q.WorkerChan <- req
}

func (q *QueueScheduler) Run() {
	q.WorkerChan = make(chan chan Request, 100)
	q.RequestChan = make(chan Request, 100)
	go func() {

		var reqQueue []Request
		var workerQueue []chan Request
		for {
			var activeReq Request
			var activeWorker chan Request

			if len(reqQueue) > 0 && len(workerQueue) > 0 {
				activeReq = reqQueue[0]
				activeWorker = workerQueue[0]
			}

			select {
			case r := <-q.RequestChan:
				reqQueue = append(reqQueue, r)
			case w := <-q.WorkerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeReq: // 把任务投递给worker协程
				workerQueue = workerQueue[1:]
				reqQueue = reqQueue[1:]
			}
		}
	}()
}
