package qsnet

import (
	"../util"
	"net"
	"../logger"
	"strconv"
)

const (
	BUFFER_SIZE = 2048
	RCHAN_SIZE  = 1024
	WCHAN_SIZE  = 1024
)

type Session struct {
	conn    net.Conn
	acceptor *Acceptor
}

func(self *Session)Close(){
	self.conn.Close()
	logger.Info("Close session:"+self.conn.RemoteAddr().String())
}

func NewSession(conn net.Conn,acceptor *Acceptor)*Session{
	return &Session{conn:conn,acceptor:acceptor}
}

type SessionManager struct {
	sessions util.QSet
}



func(self *SessionManager)AddSession(conn *Session){
	self.sessions.Add(conn)
}

func(self *SessionManager)RemoveSession(conn *Session){
	self.sessions.Remove(conn)
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