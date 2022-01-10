package p2p

import (
	"context"
	"fmt"
	"locust/protocols/generated"
	"log"
	"time"

	ggio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	noise "github.com/libp2p/go-libp2p-noise"

	ma "github.com/multiformats/go-multiaddr"
)

const clientVersion = "locust/0.0.1"

type P2PHost struct {
	host.Host
}

func NewHost() (P2PHost, error) {
	// TODO: handle errors in a cleaner way :)
	priv, _, err := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	if err != nil {
		log.Fatal(err)
		return P2PHost{}, err
	}

	listen, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", 0))
	if err != nil {
		log.Fatal(err)
		return P2PHost{}, err
	}

	host, err := libp2p.New(
		libp2p.ListenAddrs(listen),
		libp2p.Identity(priv),
		libp2p.Security(noise.ID, noise.New),
	)
	if err != nil {
		log.Fatal(err)
		return P2PHost{}, err
	}

	return P2PHost{host}, nil
}

func (n *P2PHost) SignProtoMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return n.SignData(data)
}

func (n *P2PHost) SignData(data []byte) ([]byte, error) {
	key := n.Peerstore().PrivKey(n.ID())
	return key.Sign(data)

}

func (n *P2PHost) AuthenticateMessage(message proto.Message, data *generated.MessageData) bool {
	log.Println("Authenticating message")

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

func (n *P2PHost) verifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
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

func (n *P2PHost) NewMessageData(messageId string, gossip bool) *generated.MessageData {
	nodePubKey, err := crypto.MarshalPublicKey(n.Peerstore().PubKey(n.ID()))
	if err != nil {
		panic("Failed to get public key for sender from local peer store.")
	}

	return &generated.MessageData{ClientVersion: clientVersion,
		NodeId:     peer.Encode(n.ID()),
		NodePubKey: nodePubKey,
		Timestamp:  time.Now().Unix(),
		Id:         messageId,
		Gossip:     gossip}
}

func (n *P2PHost) SendProtoMessage(id peer.ID, p protocol.ID, data proto.Message) bool {
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