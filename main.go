package main

import (
	"./sys"
	"./qsnet"
	. "./userdefine"
)

func main(){
	peer := qsnet.NewTCPPeer("0.0.0.0:8802")
	proc := NewMessageProcessor()
	peer.Start(proc)

	sys.DelayExit(func() {
		peer.Close()
		proc.Close()
	})
}
