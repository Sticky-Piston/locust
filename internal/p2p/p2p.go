package p2p

import (
	"context"
	"fmt"
	"locust/domain"

	"log"
	"os"
	"os/signal"
	"syscall"
)

type P2PImpl struct {
	Node           *Node
	ProfileUsecase domain.ProfileUsecase
}

func NewP2PProtocol(profileUsecase domain.ProfileUsecase) *P2PImpl {
	node := NewNode()

	return &P2PImpl{
		Node:           node,
		ProfileUsecase: profileUsecase,
	}
}

func (p *P2PImpl) Run(peer string, rendezvous string) {
	ctx, cancel := context.WithCancel(context.Background())

	var discoveryPeers addrList

	discoveryPeers.Set(peer)

	dht, err := NewDHT(ctx, p.Node, discoveryPeers)
	if err != nil {
		log.Fatal(err)
	}

	go Discover(ctx, p.Node, dht, rendezvous)

	log.Printf("Host ID: %s", p.Node.ID().Pretty())
	log.Printf("Connect to me on:")
	for _, addr := range p.Node.Addrs() {
		log.Printf("  %s/p2p/%s", addr, p.Node.ID().Pretty())
	}

	// Loop until cancelled
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Println("Exiting...")

	cancel()

	if err := p.Node.Close(); err != nil {
		panic(err)
	}

	os.Exit(0)
}
