package qsnet

import "net"
import "../logger"
import "../sys"

type Peer interface {
	Start(processor Processor)
}

type TCPPeer struct{
	local_addr string
	sessionManager *SessionManager
	acceptor *Acceptor
	processor Processor
}

func (self *TCPPeer)onAccept(conn net.Conn){
	s := NewSession(self,conn)
	self.sessionManager.AddSession(s)
	logger.Info("ACCEPT:"+conn.RemoteAddr().String())
	sys.StartGoroutines(func() {
		s.StartRead(self.onRead)
	})
}

func (self *TCPPeer)onClose(session *Session){
	if session.IsValid(){
		session.Close()
		self.processor.OnSessionExitEvent(CloseEvent{session})
		self.sessionManager.RemoveSession(session)
	}
}

func (self *TCPPeer)onRead(session *Session,buf []byte){
	sys.StartGoroutines(func() {
		msg := MessageEvent{Session:session,Buf:buf}
		self.processor.OnMessageEvent(msg)
	})
}

func (self *TCPPeer)Start(processor Processor){
	sys.StartGoroutines(func() {
		self.acceptor.StartAccept(self.local_addr,self.onAccept)
	})

	self.processor = processor
	sys.StartGoroutines(func() {
		processor.StartLoop()
	})

}

func (self *TCPPeer)Close(){
	self.acceptor.Close()
	self.sessionManager.Clear()
}

func NewTCPPeer(addr string)*TCPPeer{
	return &TCPPeer{
		local_addr:addr,
		sessionManager:NewSessionManager(),
		acceptor:NewAcceptor(),
	}
}
