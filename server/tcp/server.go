package tcp

import (
	"fmt"
	"strings"
	"time"

	"myutil/tcp"

	"disconf/server/dao"
	"disconf/server/model"
	tcpmodel "disconf/server/model/tcp"
	"disconf/server/util"
)

type server struct {
	registerPath map[string][]model.Client
	clients      map[string]model.Client
	server       *tcp.TCPServer
	event        chan *tcpmodel.NotifyMsg
	exit         chan struct{}
	alive        int
}

func NewServer(tcpAddr string) *server {
	srv := &server{
		registerPath: make(map[string][]model.Client),
		clients:      make(map[string]model.Client),
		event:        make(chan *tcpmodel.NotifyMsg),
		exit:         make(chan struct{}),
		alive:        60,
	}
	callback := NewCallBack(srv)
	tcpSrv := tcp.NewTCPServer(tcpAddr, callback, tcp.NewDefaultProtocol(1024))
	srv.server = tcpSrv
	return srv
}

func (t *server) Run() error {
	if err := t.server.ListenAndServe(); err != nil {
		return err
	}
	t.notifyLoop()
	return nil
}

func (t *server) NotifyEvent(event *tcpmodel.NotifyMsg) {
	t.event <- event
}

func (t *server) pushLatestConfig(paths []string, client model.Client) {
	for _, path := range paths {
		value, err := parsePath(path)
		if err != nil {
			util.G_log.Error("[parse path fail] [path:%s] [client:%s]", path, client.GetName())
			return
		}
		appName := value[0]
		clusterName := value[1]
		namespaceName := value[2]
		release, err := dao.G_release.GetLatestRelease(appName, clusterName, namespaceName)
		if err != nil {
			return
		}
		msg := &tcpmodel.NotifyMsg{
			Format:    release.Format,
			Namespace: release.NamespaceName,
			Value:     release.Value,
		}
		client.Notify(msg)
	}
}

func (t *server) notifyLoop() {
	util.G_log.Info("[notify loop start%s]", "...")
	go func() {
		for {
			select {
			case msg := <-t.event:
				util.G_log.Info("[handle notify event] [msg:%+v]", *msg)
				t.notify(msg)
			case <-t.exit:
				break
			}
		}
	}()
}

func (t *server) notify(msg *tcpmodel.NotifyMsg) {
	clients, ok := t.registerPath[msg.Namespace]
	if !ok {
		return
	}
	for _, client := range clients {
		if t.isClientAlive(client) {
			client.Notify(msg)
		} else {
			client.Close()
		}
	}
}

func (t *server) isClientAlive(client model.Client) bool {
	now := time.Now().Unix()
	if (now - client.GetHeartBeat().Unix()) > int64(t.alive) {
		return false
	}
	return true
}

func parsePath(path string) ([]string, error) {
	paths := strings.Split(path, "/")
	if len(paths) != 3 {
		return nil, fmt.Errorf("register message is error, %s", path)
	}
	return paths, nil
}
