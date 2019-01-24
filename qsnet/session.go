package qsnet

const (
	BUFFER_SIZE = 2048
)

type Session interface{
	Start()
	StartRead()
	Close()
	Send(buf []byte)
	IsValid()bool
	RemoteAddr() string
}

