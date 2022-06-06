/*
Copyright © 2022 tchiunam

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	goflag "flag"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cmdaws "github.com/tchiunam/axolgo-cli/pkg/cmd/aws"
	"k8s.io/klog/v2"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "axolgo",
	Short: "Axolgo is Axolotl in Golang",
	Long: `Axolgo is the Golang series of Axolotl. It is a
package of a variety of tools and libraries.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialize flag's default flagset
	klog.InitFlags(nil)

	// Add klog flags to cobra
	goFlagSet := goflag.NewFlagSet("", goflag.PanicOnError)
	klog.InitFlags(goFlagSet)
	rootCmd.PersistentFlags().AddGoFlagSet(goFlagSet)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.axolgo.yaml)")

	configureCommandStructure()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		klog.Errorf("Failed to bind flags to viper: %v", err)
	}
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".axolgo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".axolgo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		klog.InfoS("Using config", "file", viper.ConfigFileUsed())
	}

	// Get verbositgy from viper
	if viper.GetString("v") == "0" {
		if err := goflag.Set("v", viper.GetString("v")); err != nil {
			klog.Errorf("%v", err)
			return
		}
	}
}

func configureCommandStructure() {
	rootCmd.AddCommand(cmdaws.AwsCmd)
}
