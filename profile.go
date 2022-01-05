package main

import (
	"locust/pb"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

const profileRequest = "/profile/profilereq/0.0.1"
const profileResponse = "/profile/profileresp/0.0.1"

type ProfileProtocol struct {
	node     *Node
	requests map[string]*pb.ProfileRequest
	done     chan bool
}

func NewProfileProtocol(node *Node) *ProfileProtocol {
	p := &ProfileProtocol{node: node, requests: make(map[string]*pb.ProfileRequest)}

	node.SetStreamHandler(profileRequest, p.onProfileRequest)
	node.SetStreamHandler(profileResponse, p.onProfileResponse)

	return p
}

func (p *ProfileProtocol) onProfileRequest(s network.Stream) {

}

func (p *ProfileProtocol) onProfileResponse(s network.Stream) {

}

func (p *ProfileProtocol) Broadcast(peers []peer.AddrInfo) bool {
	return true
}
