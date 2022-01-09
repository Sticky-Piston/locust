package core

import (
	"fmt"
	"io/ioutil"
	"locust/protocols"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

const profileRequest = "/profile/profilereq/0.0.1"
const profileResponse = "/profile/profileresp/0.0.1"

type ProfileProtocol struct {
	node     *Node
	requests map[string]*protocols.ProfileRequest
}

func NewProfileProtocol(node *Node) *ProfileProtocol {
	p := &ProfileProtocol{node: node, requests: make(map[string]*protocols.ProfileRequest)}

	node.SetStreamHandler(profileRequest, p.OnProfileRequest)
	node.SetStreamHandler(profileResponse, p.OnProfileResponse)

	return p
}

func (p *ProfileProtocol) OnProfileRequest(s network.Stream) {
	log.Println("Executing profile request")

	data := &protocols.ProfileRequest{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		s.Reset()
		log.Println(err)
		return
	}

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%s: Received profile request from %s. Message: %s", s.Conn().LocalPeer(), s.Conn().RemotePeer(), data.Message)

	valid := p.node.AuthenticateMessage(data, data.MessageData)
	if !valid {
		log.Println("Failed to authenticate message")
		return
	}

	// TODO IMPLEMENT PROFILE STORAGE
	profile := &protocols.ProfileResponse{
		MessageData: p.node.NewMessageData(uuid.New().String(), false),
		Title:       "Test titel",
		Summary:     "Test summary",
	}

	ok := p.node.SendProtoMessage(s.Conn().RemotePeer(), profileResponse, profile)
	if !ok {
		log.Println("Failed to send message")
		return
	}

	log.Println("Send profile to:", s.Conn().RemotePeer())
}

func (p *ProfileProtocol) OnProfileResponse(s network.Stream) {
	log.Println("Executing profile response")
	// TODO
}

func (p *ProfileProtocol) GetProfileFromPeer(peer *peer.AddrInfo) (*protocols.ProfileResponse, error) {
	req := &protocols.ProfileRequest{
		MessageData: p.node.NewMessageData(uuid.New().String(), false),
		Message:     fmt.Sprintf("Profile request from %s", p.node.ID()),
	}

	// sign the data
	signature, err := p.node.SignProtoMessage(req)
	if err != nil {
		log.Println("failed to sign pb data")
		return nil, err
	}

	// add the signature to the message
	req.MessageData.Sign = signature

	//log.Println(req.MessageData.Sign)

	ok := p.node.SendProtoMessage(peer.ID, profileRequest, req)
	if !ok {
		return nil, nil
	}

	// store ref request so response handler has access to it
	//p.requests[req.MessageData.Id] = req
	log.Printf("%s: Profile request to: %s was sent. Message Id: %s, Message: %s", p.node.ID(), peer.ID, req.MessageData.Id, req.Message)

	time.Sleep(30 * time.Second)

	return &protocols.ProfileResponse{}, nil
}
