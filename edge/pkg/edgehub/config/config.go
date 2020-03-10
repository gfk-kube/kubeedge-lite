package config

import (
	"strings"
	"sync"

	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edgecore/v1alpha1"
)

const (
	handshakeTimeoutDefault = 60
	readDeadlineDefault     = 15
	writeDeadlineDefault    = 15

	heartbeatDefault = 15

	protocolDefault   = protocolWebsocket
	protocolWebsocket = "websocket"
	protocolQuic      = "quic"
)

var Config Configure
var once sync.Once

type Configure struct {
	v1alpha1.EdgeHub
	WebSocketURL string
	NodeName     string
}

func InitConfigure(eh *v1alpha1.EdgeHub, nodeName string) {
	once.Do(func() {
		Config = Configure{
			EdgeHub:      *eh,
			WebSocketURL: strings.Join([]string{"wss:/", eh.WebSocket.Server, eh.ProjectID, nodeName, "events"}, "/"),
			NodeName:     nodeName,
		}
	})
}
