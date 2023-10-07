package igame

type OnMessage func(msgType int, msgInfo []byte, id string)
type OnServerEvent func(msgType int, eventData []byte)

type Module interface {
	New()
	Init()
	Base()
	BeginPlay()
	UpLogic()
	UpPlayer()
	UnLoad()
	GetMsgHandleFunc(msgType int) OnMessage
	GetServerEvent() map[int]OnServerEvent
	GetModelName() string
	GetCoroutineGroup() int
}
