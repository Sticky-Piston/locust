package server

import (
	"fmt"
	"locust/handlers"
	"locust/internal/helpers"
	"locust/internal/p2p"
	"log"

	"github.com/libp2p/go-libp2p-core/protocol"
)

type Node struct {
	*p2p.P2PHost
}

func NewNode(host *p2p.P2PHost) *Node {
	node := &Node{P2PHost: host}

	profileHandler := handlers.NewProfileHandler(node.P2PHost)
	node.RegisterHandler("profile", profileHandler)

	return node
}

func (n *Node) RegisterHandler(namespace string, handler handlers.Handler) {
	log.Println(fmt.Sprintf("Registering handlers for namespace: %s", namespace))

	protocolRequest := helpers.GenerateProtocolIDRequest(namespace, handler.Version())
	protocolResponse := helpers.GenerateProtocolIDResponse(namespace, handler.Version())

	n.SetStreamHandler(protocol.ID(protocolRequest), handler.Request)
	n.SetStreamHandler(protocol.ID(protocolResponse), handler.Response)
}
