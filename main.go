package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/urfave/cli"

	"github.com/multiformats/go-multiaddr"
	ma "github.com/multiformats/go-multiaddr"
)

type Config struct {
	Port       int
	ProtocolID string
	Rendezvous string
	Seed       int64

	Node bool
}

func main() {

	var discoveryPeers addrList

	app := &cli.App{
		Name:  "Locust",
		Usage: "Run a locust node or communicate with the swarm",
		Action: func(c *cli.Context) error {
			if c.Bool("node") {
				done := make(chan bool, 1)

				host := makeNode(done)

				log.Printf("Host ID: %s", host.ID().Pretty())
				log.Printf("Connect to me on:")
				for _, addr := range host.Addrs() {
					log.Printf("  %s/p2p/%s", addr, host.ID().Pretty())
				}

				ctx, cancel := context.WithCancel(context.Background())

				discoveryPeers.Set(c.String("peer"))

				dht, err := NewDHT(ctx, host, discoveryPeers)
				if err != nil {
					log.Fatal(err)
				}

				go Discover(ctx, host, dht, c.String("rendezvous"))

				run(host, cancel)
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "node",
				Usage: "Run a locust node",
			},
			&cli.StringFlag{
				Name:  "rendezvous",
				Value: "pwnintended/locust",
				Usage: "rendezvous string",
			},
			&cli.StringFlag{
				Name:  "peer",
				Value: "",
				Usage: "List of peer addresses",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(h host.Host, cancel func()) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Println("Exiting...")

	cancel()

	if err := h.Close(); err != nil {
		panic(err)
	}

	os.Exit(0)
}

func makeNode(done chan bool) *Node {
	// TODO: handle errors :)

	priv, _, _ := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	listen, _ := ma.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", 0))
	host, _ := libp2p.New(
		libp2p.ListenAddrs(listen),
		libp2p.Identity(priv),
	)

	return NewNode(host, done)

}

type addrList []multiaddr.Multiaddr

func (al *addrList) String() string {
	strs := make([]string, len(*al))
	for i, addr := range *al {
		strs[i] = addr.String()
	}
	return strings.Join(strs, ",")
}

func (al *addrList) Set(value string) error {
	addr, err := multiaddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*al = append(*al, addr)
	return nil
}
