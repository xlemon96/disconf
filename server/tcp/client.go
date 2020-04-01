package tcp

import (
	"encoding/json"
	"time"

	"myutil/tcp"

	tcpmodel "disconf/server/model/tcp"
)

type client struct {
	heartBeatTime time.Time
	conn          *tcp.Conn
}

func NewClient(heartBeatTime time.Time, conn *tcp.Conn) *client {
	return &client{
		heartBeatTime: heartBeatTime,
		conn:          conn,
	}
}

func (t *client) Notify(msg *tcpmodel.NotifyMsg) error {
	bMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	packet := &tcp.DefaultPacket{
		Type: SendConfig,
		Body: bMsg,
	}
	t.conn.WritePacket(packet)
	return nil
}

func (t *client) GetName() string {
	return t.conn.GetRemoteAddr()
}

func (t *client) GetHeartBeat() time.Time {
	return t.heartBeatTime
}

func (t *client) UpdateHeartBeat() {
	t.heartBeatTime = time.Now()
}

func (t *client) Close() {
	t.conn.Close()
}
