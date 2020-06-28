package scheduler

// 调度器
type Scheduler interface {
	Submit(Request)
	ConfigureWorkChan(chan Request)
}

// 简单的调度器
type SimpleScheduler struct {
	workerChan chan Request
}

// 提交请求任务到调度器
func (s *SimpleScheduler) Submit(request Request) {
	go func() {
		s.workerChan<-request
	}()
}

// 设置调度器Channel
func (s *SimpleScheduler) ConfigureWorkChan(reqChan chan Request) {
	s.workerChan = reqChan
}
