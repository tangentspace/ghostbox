// Copyright © 2018 Keith Clawson <keith@tangentspace.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{Use: "boxctl"}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a service to help keep your app up",
	Long:  `A service that interacts with Nomad to automate operation of your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Starting server")
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.ghostbox.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ghostbox" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ghostbox")
	}

	// Allow using environment variables along with flags
	viper.AutomaticEnv()

	// Load config file
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Loading config file:", viper.ConfigFileUsed())
	}
}

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
