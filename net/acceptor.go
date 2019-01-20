package net

import "net"

type Acceptor struct {
	sessionManager *SessionManager
	listener net.Listener
}

func NewAcceptor()*Acceptor{
	return &Acceptor{sessionManager:NewSessionManager()}
}
