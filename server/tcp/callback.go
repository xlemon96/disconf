package tcp

import (
	"time"

	"myutil/tcp"

	"disconf/server/model"
	tcpmodel "disconf/server/model/tcp"
	"disconf/server/util"
)

type callback struct {
	server *server
}

func NewCallBack(server *server) *callback {
	return &callback{
		server: server,
	}
}

func (t *callback) OnConnected(conn *tcp.Conn) {
	util.G_log.Info("receive new connect")
}

func (t *callback) OnDisconnected(conn *tcp.Conn) {
}

func (t *callback) OnError(err error) {
}

func (t *callback) OnMessage(conn *tcp.Conn, packet tcp.Packet) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	pkt := packet.(*tcp.DefaultPacket)
	switch pkt.Type {
	case Register:
		register := &tcpmodel.Register{}
		register.Decode(pkt.Body)
		client := NewClient(time.Now(), conn)
		util.G_log.Info("[receive new register agent] [ip:%s] [registerinfo:%+v]", client.GetName(), *register)
		_, ok := t.server.clients[client.GetName()]
		if !ok {
			t.server.clients[client.GetName()] = client
		}
		for _, path := range register.Namespace {
			item, ok := t.server.registerPath[path]
			if !ok {
				item = make([]model.Client, 0)
				t.server.registerPath[path] = item
			}
			t.server.registerPath[path] = append(t.server.registerPath[path], client)
		}
		util.G_log.Info("[start push latest config] [ip:%s]", client.GetName())
		t.server.pushLatestConfig(register.Namespace, client)
	case HeartBeat:
		client, ok := t.server.clients[conn.GetRemoteAddr()]
		if ok {
			client.UpdateHeartBeat()
		}
	}
}
