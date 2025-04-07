package streamclient

import "github.com/winc-link/hummingbird/common/dtos"

type StreamClient interface {
	Send(data dtos.RpcData)
	Recv() <-chan dtos.RpcData
}
