package scheduler

// 调度器
type Scheduler interface {
	Submit(Request)
	ConfigureWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

// 简单的调度器
type SimpleScheduler struct {
	workerChan chan Request
}

func (s *SimpleScheduler) WorkerReady(requests chan Request) {
	panic("implement me")
}

func (s *SimpleScheduler) Run() {
	panic("implement me")
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
