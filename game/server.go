package game

import (
	"encoding/json"
	"github.com/nathonNot/go-gdt/config"
	"github.com/nathonNot/go-gdt/igame"
	"github.com/nathonNot/go-gdt/log"
	"sync"
	"time"
)

type Server struct {
	GameModules        map[string]igame.Module
	GameModulesRunning map[int][]igame.Module // 运行时，协程组id：模块数组
	GameEventMap       map[int][]igame.OnServerEvent
	InServerEventMap   map[int][]igame.OnInServerEvent
	PlayerManage       sync.Map // 玩家管理
	Connect2Player     sync.Map // 连接到玩家的映射
	Running            bool
}

var instance Server

func GetGameServerInstance() *Server {
	return &instance
}

func DispatchEventInServer(msgType int, eventData interface{}) {
	d, _ := json.Marshal(eventData)
	instance.DispatchEventInServer(msgType, d)
}

func (gameServer *Server) DispatchEventInServer(msgType int, eventData interface{}) {
	funcArr, ok := gameServer.InServerEventMap[msgType]
	if !ok {
		//log.ServerLog().Error("not find gamer server event", msgType)
		return
	}
	for _, f := range funcArr {
		f(msgType, eventData)
	}
}

func ServerInit(mode []igame.Module) {

	instance = Server{}
	instance.GameEventMap = make(map[int][]igame.OnServerEvent)
	instance.InServerEventMap = make(map[int][]igame.OnInServerEvent)
	instance.GameModules = make(map[string]igame.Module)
	instance.GameModulesRunning = make(map[int][]igame.Module)
	instance.GameModuleStart(mode)
	instance.Running = true
}

func (gameServer *Server) GameModuleStart(mode []igame.Module) {
	for _, m := range mode {
		m.New()
		m.Init()
		m.Base()
		gameServer.GameModules[m.GetModelName()] = m
		groupId := m.GetCoroutineGroup()
		gameServer.GameModulesRunning[groupId] = append(gameServer.GameModulesRunning[groupId], m)
		for eventType, eventFunc := range m.GetServerEvent() {
			funcArr, ok := gameServer.GameEventMap[eventType]
			if ok {
				gameServer.GameEventMap[eventType] = append(funcArr, eventFunc)
			} else {
				gameServer.GameEventMap[eventType] = []igame.OnServerEvent{eventFunc}
			}
		}
		for eType, eFunc := range m.GetInServerEvent() {
			funcArr, ok := gameServer.InServerEventMap[eType]
			if ok {
				gameServer.InServerEventMap[eType] = append(funcArr, eFunc)
			} else {
				gameServer.InServerEventMap[eType] = []igame.OnInServerEvent{eFunc}
			}
		}
	}
	for _, m := range gameServer.GameModules {
		m.BeginPlay()
	}

	var mod igame.Module
	for groupId, v := range gameServer.GameModulesRunning {
		var modelName []string
		for m := range v {
			mod = v[m]
			name := mod.GetModelName()
			modelName = append(modelName, name)
		}
		go StartGameModuleRunning(v)
		log.ServerLog().Infof("开始加载协程组：%d, 加载模块名：%v", groupId, modelName)
	}
	go gameServer.ModulesRun()

}

var groupChan []chan *igame.EventMessage

func StartGameModuleRunning(models []igame.Module) {
	eventChan := make(chan *igame.EventMessage, 1000)
	groupChan = append(groupChan, eventChan)
	for event := range eventChan {
		for _, model := range models {
			f := model.GetMsgHandleFunc(event.MsgType)
			if f != nil {
				f(event.MsgType, event.MsgInfo, event.Id)
			}
		}
	}
}

func AddEventMessage(event *igame.EventMessage) {
	for _, v := range groupChan {
		v <- event
	}
}

func (gameServer *Server) ModulesRun() {
	serverConfig := config.GetGlobalConfig()
	for {
		timeNow := time.Now()
		frameTime := serverConfig.GetGameFrameTime()
		for _, mod := range gameServer.GameModules {
			mod.UpLogic(frameTime)
		}
		sleepTime := frameTime - time.Since(timeNow)
		time.Sleep(sleepTime)
	}
}
