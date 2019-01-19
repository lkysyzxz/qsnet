package sys

import (
	"os"
	"sync"
	"os/signal"
	"syscall"
	"projects/qsnet/logger"
)

var signal_exit chan os.Signal

var goroutines_group *sync.WaitGroup

var one sync.Once


func GetSystemExitSignal()<-chan os.Signal{
	one.Do(func() {
		signal_exit = make(chan os.Signal)
		signal.Notify(signal_exit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	})
	return signal_exit
}

func closeSystemExitSignal(){
	close(signal_exit)
}

func DelayExit(decompose func()){
	exit := GetSystemExitSignal()
	<-exit
	closeSystemExitSignal()
	if (decompose != nil) {
		decompose()
	}

	logger.Info("Start wait")
	goroutines_group.Wait()
	logger.Info("Wait done")
	os.Exit(0)
}



func init(){
	logger.Info("Init wait group")
	goroutines_group = new(sync.WaitGroup)
}

type LoopFunc func()

func StartGoroutines(f LoopFunc){
	go func() {
		goroutines_group.Add(1)
		f()
		goroutines_group.Done()
	}()
}