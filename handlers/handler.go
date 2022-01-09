package handlers

import (
	"github.com/libp2p/go-libp2p-core/network"
)

type Handler interface {
	Request(s network.Stream)
	Response(s network.Stream)
	Version() string
}
