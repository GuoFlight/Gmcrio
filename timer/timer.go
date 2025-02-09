package timer

import (
	"Gmicro/conf"
	"Gmicro/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var InitDone = make(chan bool)
var Done = make(chan bool)
var exitFlags []*chan bool
var lock sync.Mutex
var wg sync.WaitGroup

func InitTimer() {
	logger.GLogger.Info("开始初始化周期性任务")

	// 启动周期性任务
	DemoTimer()
	go Exec(DemoTimer, conf.GConf.Timer.IntervalSec)

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
				lock.Lock()
				for _, flag := range exitFlags {
					*flag <- true
				}
				wg.Wait()
				logger.GLogger.Info("timer正常退出")
				Done <- true
				done <- true
			default:
				logger.GLogger.Warn("timer收到即将忽略的信号:", s)
			}
		}
	}()
	<-done
}

func Exec(task func(), intervalSec time.Duration) {
	wg.Add(1)
	defer wg.Done()

	ticker := time.NewTicker(intervalSec * time.Second)
	defer ticker.Stop()

	flag := make(chan bool)
	lock.Lock()
	exitFlags = append(exitFlags, &flag)
	lock.Unlock()

	for {
		select {
		case <-ticker.C: // 定时触发任务
			task()
		case <-flag: // 收到退出信号
			return
		}
	}
}

// DemoTimer 周期性任务的Demo
func DemoTimer() {
	logger.GLogger.Info("开始执行DemoTimer函数")
	// time.Sleep(10 * time.Second) // do something
	logger.GLogger.Info("执行了DemoTimer函数")
}
