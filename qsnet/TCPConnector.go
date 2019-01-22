package qsnet

import "net"
import "../logger"
import "../sys"

type TCPConnector struct{
	session Session
	processor Processor
	addr string
}

func (self *TCPConnector)Start()Peer{
	conn,err := net.Dial("tcp",self.addr)
	if err != nil{
		logger.Fatal(err.Error())
		return nil
	}
	self.session = NewTCPSession(self,self.processor,conn)

	sys.StartGoroutines(func() {
		self.session.StartRead()
	})
	return self
}

func (self *TCPConnector)Stop(){
	if self.session.IsValid(){
		self.session.Close()
	}
}

func (self *TCPConnector)SetProcessor(processor Processor){
	self.processor = processor
}

func (self *TCPConnector)OnClose(session Session){
	self.Stop()
}

func NewTCPConnector(addr string)*TCPConnector{
	return &TCPConnector{addr:addr}
}
