/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	peerString string
	rendezvous string
	database   string
	configDir  string
	seed       string
	port       int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "locust",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.locust.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&configDir, "config", "", "config directory (default is $HOME/.locust/)")
	rootCmd.PersistentFlags().StringVar(&seed, "seed", "", "seed to generate keypair with")
	rootCmd.PersistentFlags().IntVar(&port, "port", 0, "port to listen on")
}

func initConfig() {

	if configDir == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		configDir = filepath.Join(home, ".locust")
	}

	// Ensure the config directory exists
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0700) // create the directory
	}

	//path := filepath.Join(home, ".locust/")

	// if _, err := os.Stat(filepath.Join(path, "config.yaml")); os.IsExist(err) {
	// 	viper.SetConfigFile(filepath.Join(path, "config.yaml"))
	// } else {

	// 	viper.Set("debug", true)

	// 	// Ensure the config directory exists
	// 	if _, err := os.Stat(path); os.IsNotExist(err) {
	// 		os.MkdirAll(path, 0700) // create the directory
	// 	}

	// 	viper.AddConfigPath(filepath.Join(home, ".locust/"))
	// 	viper.SetConfigType("yaml")
	// 	viper.SetConfigName("config")

	// 	viper.AutomaticEnv()

	// 	err := viper.SafeWriteConfig()
	// 	if err != nil {
	// 		log.Println("Config already exists")
	// 	}

	// 	//viper.SafeWriteConfigAs(filepath.Join(home, ".locust", "config.yaml"))

	// }

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Using config file: ", viper.ConfigFileUsed())

		log.Println("config result: ", viper.GetBool("debug"))

	}
}
