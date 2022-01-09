package core

import (
	"context"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func Discover(ctx context.Context, node *Node, dht *discovery.RoutingDiscovery, rendezvous string) {
	var routingDiscovery = discovery.NewRoutingDiscovery(dht)
	discovery.Advertise(ctx, routingDiscovery, rendezvous)

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			peers, err := discovery.FindPeers(ctx, routingDiscovery, rendezvous)
			if err != nil {
				log.Fatal(err)
			}

			for _, peer := range peers {
				if peer.ID == node.ID() {
					continue
				}

				if node.Network().Connectedness(peer.ID) != network.Connected {
					_, err = node.Network().DialPeer(ctx, peer.ID)
					if err != nil {
						node.Peerstore().RemovePeer(peer.ID)
						continue
					}
				}
			}
		}

	}
}
