package qsnet

import "net"
import "../sys"
import "../logger"

type TCPServer struct{
	addr string
	sessionManager *SessionManager
	acceptor *Acceptor
	processor Processor
}

func (self *TCPServer)onAccept(conn net.Conn){
	s := NewTCPSession(self,self.processor,conn)
	self.sessionManager.AddSession(s)
	logger.Info("ACCEPT:"+conn.RemoteAddr().String())
	s.Start()
}

func (self *TCPServer)OnClose(session Session){
	if session.IsValid(){
		session.Close()
		self.processor.OnSessionExitEvent(CloseEvent{session})
		self.sessionManager.RemoveSession(session)
	}
}

func (self *TCPServer)Start()Peer{
	sys.StartGoroutines(func() {
		self.acceptor.StartAccept(self.addr,self.onAccept)
	})
	return self
}

func (self *TCPServer)Stop(){
	self.acceptor.Close()
	self.sessionManager.Clear()
}

func (self *TCPServer)Close(){
	self.acceptor.Close()
	self.sessionManager.Clear()
}

func (self *TCPServer)SetProcessor(processor Processor){
	self.processor = processor
}

func NewTCPServer(addr string)*TCPServer{
	return &TCPServer{
		addr:addr,
		sessionManager:NewSessionManager(),
		acceptor:NewAcceptor(),
	}
}