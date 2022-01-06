/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"locust/internal/p2p"
	"locust/src/profile/repository/badgerRepository"
	"locust/src/profile/usecase"
	"log"

	"github.com/dgraph-io/badger"
	"github.com/spf13/cobra"
)

var peer string
var rendezvous string
var database string

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
		// Open the Badger database located in the /tmp/badger directory.
		// It will be created if it doesn't exist.
		db, err := badger.Open(badger.DefaultOptions(database))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		profileRepository := badgerRepository.NewProfileRepository(db)
		profileUsecase := usecase.NewProfileUsecase(profileRepository)

		p2p.NewP2PProtocol(profileUsecase).Run(peer, rendezvous)
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringVarP(&peer, "peer", "p", "", "Peer to connect to")
	nodeCmd.Flags().StringVarP(&rendezvous, "rendezvous", "r", "locust/network", "Rendezvous string")
	nodeCmd.Flags().StringVarP(&database, "database", "d", "/tmp/locust", "Badger database file location")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
