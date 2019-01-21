package qsnet

import (
	"../util"
	"net"
	"../logger"
	"strconv"
	"sync"
)

const (
	BUFFER_SIZE = 2048
	RCHAN_SIZE  = 1024
	WCHAN_SIZE  = 1024
)

type OnReadFunc func(session *Session,buf []byte)

type Session struct {
	peer *TCPPeer
	conn    net.Conn
	close_once sync.Once
	closed	bool
}

func(self *Session)StartRead(onReadFunc OnReadFunc){
	for{
		b := make([]byte, BUFFER_SIZE)
		n, err := self.conn.Read(b) //	可引发连接异常
		if n <= 0 || err != nil {
			logger.Fatal("network exception from "+self.conn.RemoteAddr().String()+","+err.Error())
			self.onClose()
			break
		}
		go onReadFunc(self,b[0:n])
	}
}

func (self *Session)Send(buf []byte){
	self.conn.Write(buf)
}

func(self *Session)RemoteAddr()string{
	return self.conn.RemoteAddr().String()
}

func(self *Session)Close(){
	self.close_once.Do(func() {
		self.conn.Close()
		self.closed = true
		logger.Info("Close session:"+self.conn.RemoteAddr().String())
	})
}

func (self *Session)IsValid()bool{
	return !self.closed
}

func(self *Session)onClose(){
	self.peer.onClose(self)
}

func NewSession(peer *TCPPeer,conn net.Conn)*Session{
	return &Session{peer:peer,conn:conn}
}



//	SessionManager
type SessionManager struct {
	sessions util.QSet
}



func(self *SessionManager)AddSession(session *Session){
	self.sessions.Add(session)
}

func(self *SessionManager)RemoveSession(session *Session){
	self.sessions.Remove(session)
}

func(self *SessionManager)Clear(){
	self.sessions.Range(func(i interface{}) {
		s := i.(*Session)
		s.Close()
	})
	self.sessions.Clear()
	logger.Info("Clear SessionManager: Remain "+strconv.Itoa(self.sessions.Length()))
}

func NewSessionManager()*SessionManager{
	return &SessionManager{util.NewSet()}
}