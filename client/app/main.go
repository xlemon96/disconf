package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"disconf/client"
	tcpmodel "disconf/server/model/tcp"
	"myutil/tcp"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6789")
	if err != nil {
		panic(err)
	}
	packet := &tcp.DefaultPacket{
		Type: client.Register,
	}
	register := &tcpmodel.Register{
		Namespace:[]string{"test_app/test_cluster/test_namespace"},
	}
	packet.Body = register.Encode()
	len := len(packet.Bytes())
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(len))
	conn.Write(buf)
	conn.Write(packet.Bytes())

	for {
		head := make([]byte, 4)
		_, err = io.ReadFull(conn, head)
		if err != nil {
			panic(err)
		}
		//获取packet长度
		packetLength := binary.BigEndian.Uint32(head)
		//获取paket
		p := make([]byte, packetLength)
		_, err = io.ReadFull(conn, p)
		if err != nil {
			panic(err)
		}
		//解析packet
		pp := &tcp.DefaultPacket{
			Type: p[0],
			Body: p[1:],
		}
		msg := &tcpmodel.NotifyMsg{}
		json.Unmarshal(pp.Body, msg)
		fmt.Println(msg)
	}
}