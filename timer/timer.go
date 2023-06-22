package timer

import (
	"Gmicro/conf"
	"Gmicro/logger"
	"time"
)

var TimerInitDone = make(chan bool)

func InitTimer() {
	logger.GLogger.Info("开始初始化周期性任务")

	DemoTimer()
	go func() {
		for {
			DemoTimer()
			time.Sleep(time.Duration(conf.GConf.Timer.Interval) * time.Second)
		}
	}()

	// 周期性任务初始化完成
	logger.GLogger.Info("周期性任务初始化完成")
	TimerInitDone <- true
	select {}
}

// 周期性任务的Demo
func DemoTimer() {
	// do something
	logger.GLogger.Info("执行了DemoTimer函数")
}
