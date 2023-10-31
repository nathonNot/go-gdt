package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// 埋点相关标记
const (
	isActionLog = "ActionLog"
)

// 服务器行为标记
const (
	ServerActionHourRefresh         = "ServerHourRefresh"
	ServerActionDailyRefresh        = "ServerDailyRefresh"
	ServerActionWeeklyRefresh       = "ServerActionWeeklyRefresh"
	ServerActionPlayerConnectWithIp = "ServerActionPlayerConnectWithIp"
)

// 玩家相关的行为标记
const (
	PlayerLogId        = "PlayerId"
	PlayerDbLogId      = "PlayerDbLogId"
	ConnLogId          = "ConnId"
	ApiLog             = "api"
	ModuleLog          = "module"
	RoomLogId          = "RoomId"
	MsgTypeIdLog       = "MsgTypeId"
	AddRoleScore       = "AddRoleScore"
	ServerLogId        = "ServerId"
	TeamLogId          = "TeamId"
	ContainerIdLog     = "ContainerId"
	AwsTaskLog         = "AwsTaskLog"
	AwsEniLog          = "AwsEniLog"
	DockerContainerLog = "DockerContainerLog"
	ServerActionLog    = "ServerActionLog"
	WeaponNftLog       = "WeaponNFTLog"
	WeaponDbId         = "WeaponDbIdLog"
)

type Wrapper struct {
	*logrus.Entry
}

var Logger *Wrapper

func ServerLog() *Wrapper {
	return Logger
}
func (l *Wrapper) ApiLog() *Wrapper {
	return &Wrapper{l.WithFields(logrus.Fields{
		ModuleLog: ApiLog,
	})}
}
func (l *Wrapper) WithPlayerLog(playerId string) *Wrapper {
	if playerId == "" {
		return l
	}
	return &Wrapper{
		Entry: l.WithField(PlayerLogId, playerId)}
}

func (l *Wrapper) WithPlayerDbLog(playerId uint) *Wrapper {
	idStr := fmt.Sprintf("%d", playerId)
	return &Wrapper{
		Entry: l.WithField(PlayerDbLogId, idStr)}
}
