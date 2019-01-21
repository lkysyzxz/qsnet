package qsnet


type Processor interface {
	StartLoop()
	OnMessageEvent(msg MessageEvent)
	OnSessionExitEvent(msg CloseEvent)
}
