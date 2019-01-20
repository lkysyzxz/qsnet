package net

import (
	"../util"
	"net"
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

func NewSessionManager()*SessionManager{
	return &SessionManager{util.NewSet()}
}