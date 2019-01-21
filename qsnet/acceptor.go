package qsnet

import "net"
import "../logger"

const(
	DEFAULT_LIMIT = 2048
)

type OnAcceptFunc func(conn net.Conn)

type Acceptor struct {
	listener net.Listener
	connLimitChanel chan struct{}
}

func(self *Acceptor)accept(onAccept OnAcceptFunc){
	for{
		self.connLimitChanel<- struct{}{}
		conn,err := self.listener.Accept()
		if err != nil{
			break
		}
		go onAccept(conn)
	}
}

func(self *Acceptor)Close(){
	err := self.listener.Close()
	if err != nil{
		logger.Fatal(err.Error())
		return
	}
}

func(self *Acceptor)StartAccept(addr string,acceptFunc OnAcceptFunc){
	l,err := net.Listen("tcp",addr)
	if err != nil{
		logger.Fatal(err.Error())
		return
	}
	self.listener = l
	self.accept(acceptFunc)
}

func NewAcceptor()*Acceptor{
	return &Acceptor{connLimitChanel:make(chan struct{},DEFAULT_LIMIT)}
}
