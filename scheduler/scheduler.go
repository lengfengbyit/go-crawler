package scheduler

// 调度器
type Scheduler interface {
	Submit(Request)
	ConfigureWorkChan(chan Request)
	WorkerReady(chan Request)
	GetWorkerChan() chan Request
	Run()
}

// 简单的调度器
type SimpleScheduler struct {
	workerChan chan Request
}

// 所有的worker 共用一个channel队列
func (s *SimpleScheduler) GetWorkerChan() chan Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan Request) {
	// 初始化
	return
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan Request)
	return
}

// 提交请求任务到调度器
func (s *SimpleScheduler) Submit(request Request) {
	go func() {
		s.workerChan <- request
	}()
}

// 设置调度器Channel
func (s *SimpleScheduler) ConfigureWorkChan(reqChan chan Request) {
	s.workerChan = reqChan
}
