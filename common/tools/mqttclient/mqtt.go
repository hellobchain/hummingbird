package mqttclient

import (
	"context"

	"github.com/winc-link/hummingbird/common/dtos"
)

type MQTTClient interface {
	RegisterConnectCallback(dtos.ConnectHandler)
	RegisterDisconnectCallback(dtos.CallbackHandler)
	AsyncPublish(ctx context.Context, topic string, payload []byte, isSync bool)
	Close()
	GetConnectStatus() bool
}
