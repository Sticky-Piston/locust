/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"locust/core"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
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
		// Open the Badger database located in the /tmp/badger directory.
		// It will be created if it doesn't exist.
		node := core.NewNode()

		log.Printf("Host ID: %s", node.ID().Pretty())
		log.Printf("Connect to me on:")
		for _, addr := range node.Addrs() {
			log.Printf("  %s/p2p/%s", addr, node.ID().Pretty())
		}

		ctx := context.Background()

		var discoveryPeers core.AddrList
		discoveryPeers.Set(peerString)

		dht, err := core.NewDHT(ctx, node, discoveryPeers)
		if err != nil {
			log.Fatal(err)
		}

		go core.Discover(ctx, node, dht, rendezvous)

		peerAddr, err := peer.AddrInfoFromString(peerString)
		if err != nil {
			log.Fatal(err)
		}

		node.ProfileProtocol.GetProfileFromPeer(peerAddr)
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
