package server

import (
	"context"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func Discover(ctx context.Context, host host.Host, dht *discovery.RoutingDiscovery, rendezvous string) {
	var routingDiscovery = discovery.NewRoutingDiscovery(dht)
	discovery.Advertise(ctx, routingDiscovery, rendezvous)

	ticker := time.NewTicker(time.Second * 10)
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
				if peer.ID == host.ID() {
					continue
				}

				if host.Network().Connectedness(peer.ID) != network.Connected {
					_, err = host.Network().DialPeer(ctx, peer.ID)
					if err != nil {
						host.Peerstore().RemovePeer(peer.ID)
						continue
					}
				}
			}
		}
	}
}
