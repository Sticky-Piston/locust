package p2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"

	ma "github.com/multiformats/go-multiaddr"
)

type Node struct {
	host.Host
}

func NewNode() *Node {
	// Create a new host
	host := makeHost()

	//profileProtocol := handlers.NewProfileProtocol()

	// Create a new node
	node := &Node{Host: host}

	return node
}

func makeHost() host.Host {
	// TODO: handle errors :)

	priv, _, _ := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	listen, _ := ma.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", 0))
	host, _ := libp2p.New(
		libp2p.ListenAddrs(listen),
		libp2p.Identity(priv),
	)

	return host

}
