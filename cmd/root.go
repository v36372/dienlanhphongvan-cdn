// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "master",
	Short: "Some help",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Root() *cobra.Command {
	return rootCmd
}

func SetRunFunc(run func()) {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		run()
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&rootCfgFile, "rootConfig", "", "root config file, could be overrided (default is $GOPATH/src/dienlanhphongvan-cdn/config/config.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// TODO: add here

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if rootCfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(rootCfgFile)
	} else {
		envGoPath := os.Getenv("GOPATH")
		goPaths := filepath.SplitList(envGoPath)
		if len(goPaths) == 0 {
			panic("$GOPATH is not set")
		}
		for _, goPath := range goPaths {
			configDir := filepath.Join(goPath, "src", "dienlanhphongvan-cdn", "config")
			viper.AddConfigPath(configDir)
		}
		viper.SetConfigName("config")
	}
	/*
		// uncomment this if want to override config by env
		//  read in environment variables that match
		viper.SetEnvPrefix("LZ")
		viper.AutomaticEnv()
	*/

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
