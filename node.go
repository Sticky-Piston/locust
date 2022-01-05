package main

import (
	"context"
	"fmt"
	"locust/pb"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"

	ggio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
)

const clientVersion = "locust-node/0.0.1"

type Node struct {
	host.Host
	*PingProtocol
	*ProfileProtocol
}

func NewNode(host host.Host, done chan bool) *Node {
	node := &Node{Host: host}
	node.PingProtocol = NewPingProtocol(node, done)

	return node
}

func (n *Node) authenticateMessage(message proto.Message, data *pb.MessageData) bool {
	sign := data.Sign
	data.Sign = nil

	bin, err := proto.Marshal(message)
	if err != nil {
		fmt.Println(err, "failed to marshal pb message")
		return false
	}

	data.Sign = sign

	peerId, err := peer.Decode(data.NodeId)
	if err != nil {
		fmt.Println(err, "failed to decode nodeID from base58")
		return false
	}

	return n.verifyData(bin, []byte(sign), peerId, data.NodePubKey)
}

func (n *Node) signProtoMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return n.signData(data)
}

func (n *Node) signData(data []byte) ([]byte, error) {
	key := n.Peerstore().PrivKey(n.ID())
	return key.Sign(data)

}

func (n *Node) verifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
	key, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		log.Println(err, "Failed to extract key from message key data")
		return false
	}

	idFromKey, err := peer.IDFromPublicKey(key)
	if err != nil {
		log.Println("Failed to extract peer id from public key")
		return false
	}

	if idFromKey != peerId {
		log.Println(err, "Node id and provided public key mismatch")
		return false
	}

	res, err := key.Verify(data, signature)
	if err != nil {
		log.Println("Error authenticating data")
		return false
	}

	return res
}

func (n *Node) NewMessageData(messageId string, gossip bool) *pb.MessageData {
	nodePubKey, err := crypto.MarshalPublicKey(n.Peerstore().PubKey(n.ID()))
	if err != nil {
		panic("Failed to get public key for sender from local peer store.")
	}

	return &pb.MessageData{ClientVersion: clientVersion,
		NodeId:     peer.Encode(n.ID()),
		NodePubKey: nodePubKey,
		Timestamp:  time.Now().Unix(),
		Id:         messageId,
		Gossip:     gossip}
}

func (n *Node) sendProtoMessage(id peer.ID, p protocol.ID, data proto.Message) bool {
	s, err := n.NewStream(context.Background(), id, p)
	if err != nil {
		log.Println("Failed to send message:", err)
		return false
	}
	defer s.Close()

	writer := ggio.NewFullWriter(s)
	err = writer.WriteMsg(data)
	if err != nil {
		log.Println("Failed to write message:", err)
		s.Reset()
		return false
	}

	return true
}