package qsnet

import "strconv"
import (
	"../util"
	"../logger"
)

//	SessionManager
type SessionManager struct {
	sessions util.QSet
}



func(self *SessionManager)AddSession(session Session){
	self.sessions.Add(session)
}

func(self *SessionManager)RemoveSession(session Session){
	self.sessions.Remove(session)
}

func(self *SessionManager)Clear(){
	self.sessions.Range(func(i interface{}) {
		s := i.(Session)
		s.Close()
	})
	self.sessions.Clear()
	logger.Info("Clear SessionManager: Remain "+strconv.Itoa(self.sessions.Length()))
}

func NewSessionManager()*SessionManager{
	return &SessionManager{util.NewSet()}
}
