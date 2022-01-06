package main

import (
	"fmt"
	"io/ioutil"
	"locust/pb"
	"log"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
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
	log.Println("Received message!")

	data := &pb.ProfileRequest{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		s.Reset()
		log.Println(err)
		return
	}
	s.Close()

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println("Failed to unmarshal request:", err)
		return
	}

	log.Printf("%s: Received profile request from %s. Message: %s", s.Conn().LocalPeer(), s.Conn().RemotePeer(), data.Message)

	valid := p.node.authenticateMessage(data, data.MessageData)
	if !valid {
		log.Println("Failed to authenticate message")
		return
	}

	log.Println("Message is valid")

	p.onProfileResponse(s)
}

func (p *ProfileProtocol) onProfileResponse(s network.Stream) {
	log.Println("Sending response!")

	// TODO: implement
	data := &pb.ProfileResponse{}
	// buf, err := ioutil.ReadAll(s)
	// if err != nil {
	// 	s.Reset()
	// 	log.Println(err)
	// 	return
	// }
	// s.Close()

	valid := p.node.authenticateMessage(data, data.MessageData)

	if !valid {
		log.Println("Failed to authenticate message")
		return
	} else {
		log.Println("Valid signature")
	}
}

func (p *ProfileProtocol) RequestProfile(peer *peer.AddrInfo) bool {
	log.Println("Requesting profile")

	// create message data
	req := &pb.ProfileRequest{MessageData: p.node.NewMessageData(uuid.New().String(), false),
		Message: fmt.Sprintf("Ping from %s", p.node.ID())}

	// sign the data
	signature, err := p.node.signProtoMessage(req)
	if err != nil {
		log.Println("failed to sign pb data")
		return false
	}

	// add the signature to the message
	req.MessageData.Sign = signature

	ok := p.node.sendProtoMessage(peer.ID, profileRequest, req)
	if !ok {
		log.Println("Couldn't send message")
		return false
	}

	log.Println("Send message!")

	return true
}
