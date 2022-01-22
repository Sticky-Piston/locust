package rpc

import (
	"io/ioutil"
	"locust/domain"
	"locust/internal/helpers"
	"locust/internal/p2p"
	"locust/protocols/generated"
	"log"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"google.golang.org/protobuf/proto"
)

const profileVersion = "0.0.1"

type ProfileHandler struct {
	Usecase domain.ProfileUsecase
	host    *p2p.P2PHost
}

func NewProfileHandler(host *p2p.P2PHost, usecase domain.ProfileUsecase) Handler {
	return &ProfileHandler{
		Usecase: usecase,
		host:    host,
	}
}

func (h *ProfileHandler) Request(s network.Stream) {
	log.Println("executing a Profile Request")

	data := &generated.ProfileGetRequest{}
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

	profile, err := h.Usecase.GetProfile()
	if err != nil {
		log.Println(err)
		return
	}

	profileMessage := &generated.ProfileGetResponse{
		MessageData: h.host.NewMessageData(uuid.New().String(), false),
		Author:      profile.Author,
		Summary:     profile.Summary,
	}

	signature, err := h.host.SignProtoMessage(profileMessage)
	if err != nil {
		log.Println("failed to sign response")
		return
	}

	profileMessage.MessageData.Sign = signature

	ok := h.host.SendProtoMessage(s.Conn().RemotePeer(), protocol.ID(helpers.GenerateProtocolIDResponse("profile", h.Version())), profileMessage)
	if !ok {
		log.Println("Failed to send message")
		return
	}

	log.Println("Send profile request to:", s.Conn().RemotePeer())
}

func (h *ProfileHandler) Response(s network.Stream) {
	log.Println("executing a Profile Reponse")

	data := &generated.ProfileGetResponse{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		s.Reset()
		log.Println(err)
		return
	}
	s.Close()

	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Println(err)
		return
	}
	valid := h.host.AuthenticateMessage(data, data.MessageData)
	if !valid {
		log.Println("Failed to authenticate message")
		return
	}

	log.Printf("%s: Received profile response from %s. Message id:%s. Message: %s.", s.Conn().LocalPeer(), s.Conn().RemotePeer(), data.MessageData.Id, data.Summary)
}

func (h *ProfileHandler) Version() string {
	return profileVersion
}
