/*
  author='du'
  date='2020/1/28 17:32'
*/
package simplescheduler

import "golang20200117/crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigerMasterWorkerChan(c chan engine.Request) {
	s.workChan = c
}

//就做把request送进workchan就这么一件事情。
func (s *SimpleScheduler) Submit(r engine.Request) {
	//這裡有可能會卡死，所以開一個goroutine
	go func() { s.workChan <- r }()
}
