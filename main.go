package main

import (
	"./logger"
	"./sys"
	"./util"
	"fmt"
)

func main(){
	set := util.NewSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)
	fmt.Println(set.Contains(4))
	sys.DelayExit(func() {
		logger.Info("Hook")
	})
}
