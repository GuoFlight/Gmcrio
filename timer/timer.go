package timer

import (
	"Gmicro/conf"
	"Gmicro/logger"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var InitDone = make(chan bool)
var Done = make(chan bool)
var exitFlag = false
var wg sync.WaitGroup

func InitTimer() {
	logger.GLogger.Info("开始初始化周期性任务")

	// 启动周期性任务
	DemoTimer()
	go Exec(DemoTimer, &wg)

	// 周期性任务初始化完成
	logger.GLogger.Info("周期性任务初始化完成")
	InitDone <- true

	// 优雅退出
	sig := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for {
			s := <-sig
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				logger.GLogger.Info("timer收到退出信号：", s)
				exitFlag = true
				wg.Wait()
				logger.GLogger.Info("timer正常退出")
				Done <- true
				done <- true
			default:
				fmt.Println("timer收到即将忽略的信号:", s)
			}
		}
	}()
	<-done
}

func Exec(task func(), wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		task()
		time.Sleep(time.Duration(conf.GConf.Timer.Interval) * time.Second)
		if exitFlag {
			break
		}
	}
}

// DemoTimer 周期性任务的Demo
func DemoTimer() {
	// do something
	logger.GLogger.Info("执行了DemoTimer函数")
}
