package game

import (
	"encoding/json"
	"github.com/nathonNot/go-gdt/igame"
	"github.com/nathonNot/go-gdt/log"
	"sync"
)

type Server struct {
	GameModules        map[string]igame.Module
	GameModulesRunning map[int][]igame.Module // 运行时，协程组id：模块数组
	GameEventMap       map[int][]igame.OnServerEvent
	PlayerManage       sync.Map // 玩家管理
}

var instance Server

func GetGameServerInstance() *Server {
	return &instance
}

func DispatchEventInServer(msgType int, eventData interface{}) {
	d, _ := json.Marshal(eventData)
	instance.DispatchEventInServer(msgType, d)
}

func (gameServer *Server) DispatchEventInServer(msgType int, eventData []byte) {
	funcArr, ok := gameServer.GameEventMap[msgType]
	if !ok {
		//log.ServerLog().Error("not find gamer server event", msgType)
		return
	}
	for _, f := range funcArr {
		f(msgType, eventData)
	}
}

func (gameServer *Server) GameModuleStart(mode []igame.Module) {
	gameServer.GameModulesRunning = make(map[int][]igame.Module)
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
		log.ServerLog().Infof("开始加载协程组：%d, 加载模块名：%v", groupId, modelName)
		go ModulesRun(v)
	}
}
