package container

import (
	"github.com/winc-link/hummingbird/common/pkg/di"
	"github.com/winc-link/hummingbird/common/tools/streamclient"
)

var StreamClientName = di.TypeInstanceToName((*streamclient.StreamClient)(nil))

func StreamClientFrom(get di.Get) streamclient.StreamClient {
	return get(StreamClientName).(streamclient.StreamClient)
}
