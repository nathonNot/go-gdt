package igame

type IEventMq interface {
	Name() string
	Init() error
	DispatchEvent(msgType int, eventData interface{}) // 发送服务器间消息
	WaitEvent() <-chan []byte                         // 等待服务器间消息
}
