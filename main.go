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

	// connector := qsnet.NewTCPConnector("127.0.0.1:8802")
	// qsnet.BindProcessor(connector,proc)
	// connector.Start()

	sys.DelayExit(func() {
		// connector.Stop()
		peer.Stop()
		proc.Close()
		// proc.Close()
	})
}
