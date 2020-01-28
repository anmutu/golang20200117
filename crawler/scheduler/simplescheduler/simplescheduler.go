/*
  author='du'
  date='2020/1/28 17:32'
*/
package simplescheduler

import "golang20200117/crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

//这里的workChan我们自己来配置，也就是将定义好的channel赋值给我们的workChan就完了。
func (s *SimpleScheduler) ConfigerMasterWorkerChan(c chan engine.Request) {
	s.workChan = c
}

//就做把request送进workchan就这么一件事情。
//也就是scheduler里的channel里的数据送到worker里去。
func (s *SimpleScheduler) Submit(r engine.Request) {
	//這裡有可能會卡死，所以開一個goroutine
	go func() { s.workChan <- r }()
}
