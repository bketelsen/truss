// Copyright © 2018 Brian Ketelsen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var compute string
var messaging string
var name string
var inputSourceName string
var outputSinkName string
var devops bool

type Compute string

const (
	Functions Compute = "func"
	Container Compute = "container"
)

type Messaging string

const (
	RPC        Messaging = "grpc"
	ServiceBus Messaging = "servicebus"
	JSON       Messaging = "json"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "truss",
	Short: "Generate scaffolded services.",
	Long:  `A longer description here. `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.truss.yaml)")
	rootCmd.PersistentFlags().StringVar(&compute, "compute", "", "compute target - Options: "+computeOptions())
	rootCmd.PersistentFlags().StringVar(&messaging, "messaging", "", "messaging platform - Options: "+messagingOptions())
	rootCmd.PersistentFlags().StringVar(&name, "name", "", "name of the appliction")
	rootCmd.PersistentFlags().StringVar(&inputSourceName, "input-source-name", "", "name of the input source")
	rootCmd.PersistentFlags().StringVar(&outputSinkName, "output-sink-name", "", "name of the output sink")
	rootCmd.PersistentFlags().BoolVar(&devops, "devops", false, "create registry") // and associated build tasks? Devops flow? CI?
}

// initConfig reads in config file and ENV variables if set.
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

		// Search config in home directory with name ".truss" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".truss")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func computeOptions() string {
	return strings.Join([]string{string(Functions), string(Container)}, ", ")
}
func messagingOptions() string {
	return strings.Join([]string{string(RPC), string(ServiceBus), string(JSON)}, ", ")
}
