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

package types

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// TestAxolgoConfig tests the AxolgoConfig structure
func TestAxolgoConfig(t *testing.T) {
	cases := map[string]struct {
		configFilePath               string
		logLevelverbosity            int
		awsRegion                    string
		googleApplicationCredentials string
		gcpZone                      string
	}{
		"normal config file": {
			configFilePath:               "./testdata",
			logLevelverbosity:            0,
			awsRegion:                    "ap-east-1",
			googleApplicationCredentials: "~/.gcp_credentials",
			gcpZone:                      "asia-east1-a",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			viper.AddConfigPath(tc.configFilePath)
			viper.SetConfigType("yaml")
			viper.SetConfigName("axolgo")
			assert.NoError(
				t,
				viper.ReadInConfig(),
				fmt.Sprintf("Failed to read config file %s", viper.ConfigFileUsed()))
			var axolgoConfig AxolgoConfig
			assert.NoError(t, viper.Unmarshal(&axolgoConfig), "Failed to unmarshal config file")
			assert.Equal(t, tc.logLevelverbosity, axolgoConfig.Logging.LogLevelVerbosity, "Expected log level verbosity %d, got %d", tc.logLevelverbosity, axolgoConfig.Logging.LogLevelVerbosity)
			assert.Equal(t, tc.awsRegion, axolgoConfig.AWS.Region, "Expected aws region %s, got %s", tc.awsRegion, axolgoConfig.AWS.Region)
			assert.Equal(t, tc.googleApplicationCredentials, axolgoConfig.GCP.GoogleApplicationCredentials, "Expected google application credentials %s, got %s", tc.googleApplicationCredentials, axolgoConfig.GCP.GoogleApplicationCredentials)
			assert.Equal(t, tc.gcpZone, axolgoConfig.GCP.Zone, "Expected gcp zone %s, got %s", tc.gcpZone, axolgoConfig.GCP.Zone)
		})
	}
}
