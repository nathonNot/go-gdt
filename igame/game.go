package igame

import "time"

type OnMessage func(msgType int, msgInfo []byte, id string)
type OnServerEvent func(msgType int, eventData []byte)

type Module interface {
	New()
	Init()
	Base()
	BeginPlay()
	UpLogic(frameTime time.Duration)
	UpPlayer()
	UnLoad()
	GetMsgHandleFunc(msgType int) OnMessage
	GetServerEvent() map[int]OnServerEvent
	GetModelName() string
	GetCoroutineGroup() int
}
