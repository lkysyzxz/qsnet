package userdefine

import "../qsnet"
import "../logger"

const(
	PIPE_SIZE = 2048
)


type MessageProcessor struct {
	pipe chan qsnet.MessageEvent
	close chan struct{}
}

func (self *MessageProcessor)OnSessionExitEvent(msg qsnet.CloseEvent){
	logger.Info("Session closed:"+msg.Session.RemoteAddr())
}

func (self *MessageProcessor)OnMessageEvent(msg qsnet.MessageEvent){
	self.pipe<-msg
}

func (self *MessageProcessor)StartLoop(){
	for{
		select {
			case <-self.close:
				return
			case msg := <-self.pipe:
				logger.Info("Read:"+string(msg.Buf)+". From:"+msg.Session.RemoteAddr())
				msg.Session.Send(msg.Buf)
		}
	}
	close(self.pipe)
}

func (self *MessageProcessor)Close(){
	close(self.close)
}

func NewMessageProcessor()*MessageProcessor{
	return &MessageProcessor{
		pipe:make(chan qsnet.MessageEvent,PIPE_SIZE),
		close:make(chan struct{}),
	}
}