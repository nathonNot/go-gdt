package igame

import "time"

type OnMessage func(msgType int, msgInfo []byte, id string)
type OnServerEvent func(msgType int, eventData []byte)
type OnInServerEvent func(msgType int, eventData interface{})

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
	GetInServerEvent() map[int]OnInServerEvent
	GetModelName() string
	GetCoroutineGroup() int
}
