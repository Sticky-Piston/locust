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

	//host := node.Host

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

				log.Println("Found peer:", peer.ID)

				// addrs := node.Peerstore().Addrs(peer.ID)
				// if len(addrs) == 0 {
				// 	log.Println("Discovered peer with ID:", peer.ID)
				// 	node.Peerstore().AddAddrs(peer.ID, peer.Addrs, peerstore.PermanentAddrTTL)
				// } else {
				// 	log.Println("I already know peer with ID:", peer.ID)
				// }

				if node.Network().Connectedness(peer.ID) != network.Connected {
					_, err = node.Network().DialPeer(ctx, peer.ID)
					if err != nil {
						node.Peerstore().RemovePeer(peer.ID)
						continue
					}
					//log.Println("Discovered peer with ID:", peer.ID)

				}
			}
		}

	}
}
