package qsnet

type MessageEvent struct {
	Session *Session
	Buf []byte
}

type CloseEvent struct{
	Session *Session
}
