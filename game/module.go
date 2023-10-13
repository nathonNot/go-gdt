package game

import (
	"github.com/nathonNot/go-gdt/config"
	"github.com/nathonNot/go-gdt/igame"
	"time"
)

type Module struct {
	ModuleName     string                      // 模块名
	CoroutineGroup int                         // 分配的协程组
	HandleMsgMap   map[int]igame.OnMessage     // 本模块捕获的协议id
	ServerEventMap map[int]igame.OnServerEvent // 服务器内事件
}

func (md *Module) New() {
}

func (md *Module) BeginPlay() {
}

func (md *Module) Base() {
}

func (md *Module) Init() {
}

func (md *Module) UpLogic(frameTime time.Duration) {

}

func (md *Module) UpPlayer() {

}

func (md *Module) UnLoad() {

}

func (md *Module) GetCoroutineGroup() int {
	return md.CoroutineGroup
}

func (md *Module) GetModelName() string {
	return md.ModuleName
}

func (md *Module) GetMsgHandleFunc(msgType int) igame.OnMessage {
	handlerFunc, ok := md.HandleMsgMap[msgType] /*如果确定是真实的,则存在,否则不存在 */
	if ok {
		return handlerFunc
	}
	return nil
}

func (md *Module) GetServerEvent() map[int]igame.OnServerEvent {
	return md.ServerEventMap
}

func ModulesRun(modules []igame.Module) {
	var mod igame.Module
	serverConfig := config.GetGlobalConfig()
	for {
		timeNow := time.Now()
		frameTime := serverConfig.GetGameFrameTime()
		for i := range modules {
			mod = modules[i]
			mod.UpLogic(frameTime)
		}
		time.Sleep(frameTime - time.Since(timeNow))
	}
}
