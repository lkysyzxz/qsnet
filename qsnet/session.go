package qsnet

const (
	BUFFER_SIZE = 2048
	RCHAN_SIZE  = 1024
	WCHAN_SIZE  = 1024
)

type Session interface{
	StartRead()
	Close()
	Send(buf []byte)
	IsValid()bool
	RemoteAddr() string
}

