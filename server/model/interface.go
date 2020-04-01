package model

import (
	"time"

	"disconf/server/model/tcp"
)

type Client interface {
	Notify(msg *tcp.NotifyMsg) error
	GetName() string
	GetHeartBeat() time.Time
	UpdateHeartBeat()
	Close()
}

type Server interface {
	NotifyEvent(event *tcp.NotifyMsg)
}
