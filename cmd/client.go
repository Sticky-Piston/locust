/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"locust/internal/helpers"
	"locust/internal/p2p"
	"locust/protocols/generated"
	"locust/server"
	"locust/service/profile"
	bleveProfileRepository "locust/service/profile/repository/bleve"
	"log"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// open a new index
		mapping := bleve.NewIndexMapping()
		index, err := bleve.Open(database)
		if err != nil {
			bleve.New(database, mapping)
		}

		if index == nil {
			log.Fatal("index is nil")
		}

		bleveProfileRepository := bleveProfileRepository.NewBleveProfileRepository(index)
		profileUsecase := profile.NewProfileUsecase(bleveProfileRepository)

		host, err := p2p.NewHost("blaat", port)
		if err != nil {
			log.Fatal(err)
			return
		}

		node := server.NewNode(&host, profileUsecase)

		ctx := context.Background()
		//defer cancel()

		var discoveryPeers helpers.AddrList
		discoveryPeers.Set(peerString)

		dht, err := server.NewDHT(ctx, host, discoveryPeers)
		if err != nil {
			log.Fatal(err)
		}

		go server.Discover(ctx, host, dht, rendezvous)

		peer, err := peer.AddrInfoFromString(peerString)
		if err != nil {
			log.Fatal(err)
			return
		}

		req := &generated.ProfileGetRequest{
			MessageData: host.NewMessageData(uuid.New().String(), false),
			Message:     fmt.Sprintf("Profile request from %s", node.ID()),
		}

		// sign the data
		signature, err := host.SignProtoMessage(req)
		if err != nil {
			log.Println("failed to sign pb data")
			return
		}

		// add the signature to the message
		req.MessageData.Sign = signature

		ok := host.SendProtoMessage(peer.ID, protocol.ID(helpers.GenerateProtocolIDRequest("profile", "0.0.1")), req)
		if !ok {
			return
		}

		//log.Printf("%s: Profile request to: %s was sent. Message Id: %s, Message: %s", node.ID(), peer.ID, req.MessageData.Id, req.Message)

		time.Sleep(5 * time.Second)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringVarP(&peerString, "peer", "p", "", "Peer to connect to")
	clientCmd.Flags().StringVarP(&rendezvous, "rendezvous", "r", "locust/network", "Rendezvous string")
	clientCmd.Flags().StringVarP(&database, "database", "d", "/tmp/locust", "Badger database file location")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
