/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"locust/internal/helpers"
	"locust/internal/p2p"
	"locust/server"
	"locust/service/profile"
	bleveProfileRepository "locust/service/profile/repository/bleve"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/blevesearch/bleve/v2"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
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

		host, err := p2p.NewHost()
		if err != nil {
			log.Fatal(err)
			return
		}
		node := server.NewNode(&host, profileUsecase)

		log.Printf("Host ID: %s", node.ID().Pretty())
		log.Printf("Connect to me on:")
		for _, addr := range node.Addrs() {
			log.Printf("  %s/p2p/%s", addr, node.ID().Pretty())
		}

		ctx, cancel := context.WithCancel(context.Background())

		var discoveryPeers helpers.AddrList
		discoveryPeers.Set(peerString)

		dht, err := server.NewDHT(ctx, node, discoveryPeers)
		if err != nil {
			log.Println(err)
		}

		go server.Discover(ctx, node, dht, rendezvous)

		c := make(chan os.Signal, 1)

		signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
		<-c

		fmt.Println("Exiting...")

		cancel()

		if err := node.Close(); err != nil {
			panic(err)
		}

		os.Exit(0)

	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringVarP(&peerString, "peer", "p", "", "Peer to connect to")
	nodeCmd.Flags().StringVarP(&rendezvous, "rendezvous", "r", "locust/network", "Rendezvous string")
	nodeCmd.Flags().StringVarP(&database, "database", "d", "/tmp/locust.bleve", "Badger database file location")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
