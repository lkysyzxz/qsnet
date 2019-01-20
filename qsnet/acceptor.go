package qsnet

import "net"
import "../logger"
import "../sys"

const(
	DEFAULT_LIMIT = 2048
)

type Acceptor struct {
	sessionManager *SessionManager
	listener net.Listener
	connLimitChanel chan struct{}
}

func(self *Acceptor)accept(){
	for{
		self.connLimitChanel<- struct{}{}
		conn,err := self.listener.Accept()
		if err != nil{
			break
		}
		go self.onAccept(conn)
	}
}

func(self *Acceptor)onAccept(conn net.Conn){
	s := NewSession(conn,self)
	self.sessionManager.AddSession(s)
	logger.Info("ACCEPT:"+conn.RemoteAddr().String())
}

func(self *Acceptor)Close(){
	err := self.listener.Close()
	if err != nil{
		logger.Fatal(err.Error())
		return
	}
	self.sessionManager.Clear()
}

func(self *Acceptor)Start(addr string){
	l,err := net.Listen("tcp",addr)
	if err != nil{
		logger.Fatal(err.Error())
		return
	}
	self.listener = l

	sys.StartGoroutines(func() {
		self.accept()
	})
}

func NewAcceptor()*Acceptor{
	return &Acceptor{sessionManager:NewSessionManager(),connLimitChanel:make(chan struct{},DEFAULT_LIMIT)}
}
