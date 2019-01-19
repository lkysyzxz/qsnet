package main

import (
	"projects/qsnet/sys"
	"projects/qsnet/logger"
)

func main(){
	sys.DelayExit(func() {
		logger.Info("Hook")
	})
}
