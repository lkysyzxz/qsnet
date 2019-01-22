package qsnet

type Peer interface {
	Start()Peer
	Stop()
	SetProcessor(processor Processor)
	OnClose(session Session)
}

func BindProcessor(peer Peer,processor Processor){
	peer.SetProcessor(processor)
}

