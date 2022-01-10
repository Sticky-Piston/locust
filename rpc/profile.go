package rpc

import (
	"io/ioutil"
	"locust/internal/p2p"
	"locust/protocols/generated"
	"log"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/network"
	"google.golang.org/protobuf/proto"
)

const profileVersion = "0.0.1"

type ProfileHandler struct {
	host *p2p.P2PHost
}

func NewProfileHandler(host *p2p.P2PHost) Handler {
	return &ProfileHandler{
		host: host,
	}
}

func (h *ProfileHandler) Request(s network.Stream) {
	log.Println("executing a Profile Request")

	data := &generated.ProfileRequest{}
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

	valid := h.host.AuthenticateMessage(data, data.MessageData)
	if !valid {
		log.Println("Failed to authenticate message")
		return
	}

	// TODO IMPLEMENT PROFILE STORAGE
	profile := &generated.ProfileResponse{
		MessageData: h.host.NewMessageData(uuid.New().String(), false),
		Title:       "Test titel",
		Summary:     "Test summary",
	}

	ok := h.host.SendProtoMessage(s.Conn().RemotePeer(), profileVersion, profile)
	if !ok {
		log.Println("Failed to send message")
		return
	}

	log.Println("Send profile to:", s.Conn().RemotePeer())
}

func (h *ProfileHandler) Response(s network.Stream) {
	log.Println("executing a Profile Reponse")
}

func (h *ProfileHandler) Version() string {
	return profileVersion
}
