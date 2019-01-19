package main

import (
	"./sys"
	"./logger"
)

func main(){
	sys.DelayExit(func() {
		logger.Info("Hook")
	})
}
