package qsnet

import (
	"net"
	"sync"
	"../logger"
	"../sys"
)

type TCPSession struct {
	peer Peer
	processor Processor
	conn    net.Conn
	close_once sync.Once
	closed	bool
}

func(self *TCPSession)StartRead(){
	for{
		b := make([]byte, BUFFER_SIZE)
		n, err := self.conn.Read(b) //	可引发连接异常
		if n <= 0 || err != nil {
			logger.Fatal("network exception from "+self.conn.RemoteAddr().String()+","+err.Error())
			self.onClose()
			break
		}
		sys.StartGoroutines(func() {
			self.processor.OnMessageEvent(MessageEvent{self,b[0:n]})
		})
	}
}

func (self *TCPSession)Send(buf []byte){
	self.conn.Write(buf)
}

func(self *TCPSession)RemoteAddr()string{
	return self.conn.RemoteAddr().String()
}

func(self *TCPSession)Close(){
	self.close_once.Do(func() {
		self.conn.Close()
		self.closed = true
		logger.Info("Close TCPSession:"+self.conn.RemoteAddr().String())
	})
}

func (self *TCPSession)IsValid()bool{
	return !self.closed
}

func(self *TCPSession)onClose(){
	self.peer.OnClose(self)
}

func NewTCPSession(peer Peer,processor Processor,conn net.Conn)*TCPSession{
	return &TCPSession{peer:peer,processor:processor,conn:conn}
}
