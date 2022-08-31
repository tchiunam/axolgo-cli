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

package util

import (
	goflag "flag"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/viper"
	"github.com/tchiunam/axolgo-cli/pkg/types"
	"k8s.io/klog/v2"
)

func InitAxolgoConfig(cfgFilePath string) {
	viper.AutomaticEnv() // read in environment variables that match

	// Use config file path from the flag
	if cfgFilePath == "" {
		if cfgFilePath = viper.GetString("AXOLGO_CONFIG_PATH"); cfgFilePath == "" {
			cfgFilePath = "./config"
		}
	}

	viper.AddConfigPath(cfgFilePath)

	viper.SetConfigType("yaml")
	viper.SetConfigName("axolgo")
	// If base config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		// This logging level is triggered by command line argument only
		// because the configuration file has not been loaded yet
		klog.V(1).InfoS("Using base config", "file", viper.ConfigFileUsed())
	} else {
		klog.Fatal(err)
	}

	// Read multiple sets of configuration file
	for _, configSet := range []string{"aws", "gcp", "logging"} {
		// Check if the config file exists
		configName := "axolgo-" + configSet
		// Config file is optional
		if _, err := os.Stat(filepath.Join(cfgFilePath, configName+".yaml")); err == nil {
			// If a config file is found, read it in
			viper.SetConfigName(configName)
			if err := viper.MergeInConfig(); err == nil {
				// This logging level is triggered by command line argument only
				// because the configuration file has not been loaded yet
				klog.V(1).InfoS("Using "+configSet+" config", "file", viper.ConfigFileUsed())
			} else {
				klog.Fatalf("Failed to read axolgo-%v.yaml.", configSet)
			}
		}
	}

	// Parse AxolgoConfig and put it into viper
	var axolgoConfig types.AxolgoConfig
	if err := viper.Unmarshal(&axolgoConfig); err != nil {
		klog.Fatalf("Encountered error when parsing axolgo configuration file: %v", err)
	}
	viper.Set("axolgo-config", axolgoConfig)

	// Get verbosity from viper
	if axolgoConfig.Logging.LogLevelVerbosity > 0 {
		if err := goflag.Set("v", strconv.Itoa(axolgoConfig.Logging.LogLevelVerbosity)); err != nil {
			klog.Errorf("%v", err)
		}
	}
}
