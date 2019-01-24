package qsnet

import (
	"../logger"
	"../sys"
	"../util"
	"net"
	"sync"
)

const (
	CACHE_SIZE = 1024
)

type TCPSession struct {
	peer       Peer
	processor  Processor
	conn       net.Conn
	close_once sync.Once
	closed     bool
	ch_send    chan []byte
	ch_read    chan []byte
}

func (self *TCPSession) Start() {
	sys.StartGoroutines(func() {
		self.StartRead()
	})

	sys.StartGoroutines(func() {
		self.processRead()
	})

	sys.StartGoroutines(func() {
		self.processSend()
	})
}

func (self *TCPSession) StartRead() {
	for {
		b := make([]byte, BUFFER_SIZE)
		n, err := self.conn.Read(b) // 	可引发连接异常
		if n <= 0 || err != nil {
			logger.Fatal("network exception from " + self.conn.RemoteAddr().String() + "," + err.Error())
			self.onClose()
			break
		}
		self.ch_read <- b[0:n]
	}
}

func (self *TCPSession) processRead() {
	stream := util.NewStreamBuffer()
	for b := range self.ch_read {
		if b != nil {
			stream.Append(b)
			for stream.Len() > 4 {
				dlen := stream.ReadInt()
				if !stream.Empty() && dlen <= stream.Len() {
					data := stream.ReadNBytes(dlen)
					self.processor.OnMessageEvent(MessageEvent{self, data})
				} else {
					stream.Undo()
					break
				}
			}
		} else {
			break
		}
	}
}

func (self *TCPSession) Send(buf []byte) {
	if self.IsValid() {
		self.ch_send <- buf
	}
}

func (self *TCPSession) RemoteAddr() string {
	return self.conn.RemoteAddr().String()
}

func (self *TCPSession) Close() {
	self.close_once.Do(func() {
		self.conn.Close()
		close(self.ch_read)
		close(self.ch_send)
		self.closed = true
		logger.Info("Close TCPSession:" + self.conn.RemoteAddr().String())
	})
}

func (self *TCPSession) IsValid() bool {
	return !self.closed
}

func (self *TCPSession) onClose() {
	self.peer.OnClose(self)
}

func (self *TCPSession) processSend() {
	for b := range self.ch_send {
		if b != nil {
			stream := util.NewStreamBuffer()
			stream.WriteInt(len(b))
			stream.Append(b)
			n, err := self.conn.Write(stream.Bytes())
			if n <= 0 || err != nil {
				logger.Fatal(err.Error())
				panic(err)
				return
			}
		} else {
			break
		}
	}
}

func NewTCPSession(peer Peer, processor Processor, conn net.Conn) *TCPSession {
	return &TCPSession{
		peer:      peer,
		processor: processor,
		conn:      conn,
		ch_send:   make(chan []byte, CACHE_SIZE),
		ch_read:   make(chan []byte, CACHE_SIZE),
	}
}
