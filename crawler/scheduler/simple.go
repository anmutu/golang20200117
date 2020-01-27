/*
  author='du'
  date='2020/1/26 22:22'
*/
package scheduler

import "golang20200117/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigerMsaterWorkerChan(c chan engine.Request) {
	//panic("implement me")
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	s.workerChan <- r
}
