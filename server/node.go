package server

import (
	"fmt"
	"locust/domain"
	"locust/internal/helpers"
	"locust/internal/p2p"
	"locust/rpc"
	"log"

	"github.com/libp2p/go-libp2p-core/protocol"
)

type Node struct {
	*p2p.P2PHost
}

func NewNode(host *p2p.P2PHost, profileUseCase domain.ProfileUsecase) *Node {
	node := &Node{P2PHost: host}

	profileHandler := rpc.NewProfileHandler(node.P2PHost, profileUseCase)
	node.RegisterHandler("profile", profileHandler)

	return node
}

func (n *Node) RegisterHandler(namespace string, handler rpc.Handler) {
	log.Println(fmt.Sprintf("Registering handlers for namespace: %s", namespace))

	protocolRequest := helpers.GenerateProtocolIDRequest(namespace, handler.Version())
	protocolResponse := helpers.GenerateProtocolIDResponse(namespace, handler.Version())

	n.SetStreamHandler(protocol.ID(protocolRequest), handler.Request)
	n.SetStreamHandler(protocol.ID(protocolResponse), handler.Response)
}
