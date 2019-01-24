package main

import (
	"./qsnet"
	"./sys"
	"./userdefine"
)

func main(){
	peer := qsnet.NewTCPServer("0.0.0.0:8802")
	proc := userdefine.NewMessageProcessor()
	qsnet.BindProcessor(peer,proc)
	peer.Start()
	proc.StartLoop()

	sys.DelayExit(func() {
		peer.Stop()
		proc.Close()
	})
}
