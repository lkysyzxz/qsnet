package main

import (
	"./qsnet"
	"./sys"
)

func main(){
	acceptor := qsnet.NewAcceptor()
	acceptor.Start("0.0.0.0:8802")
	sys.DelayExit(func() {
		acceptor.Close()
	})
}
