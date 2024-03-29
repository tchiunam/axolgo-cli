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
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/tchiunam/axolgo-cli/pkg/types"
)

// TestInitAxolgoConfig calls InitAxolgoConfig to make sure it initialize a valid configuration.
func TestInitAxolgoConfig(t *testing.T) {
	cfgFilePath := filepath.Join(filepath.Dir(""), "..", "testdata", "config")

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{
		"cmd",
		"-v", "3"}

	t.Run("init with parameter", func(t *testing.T) {
		assert.NotPanics(t, func() { InitAxolgoConfig(cfgFilePath) })
	})

	os.Setenv("AXOLGO_CONFIG_PATH", cfgFilePath)

	axolgoConfig := viper.Get("axolgo-config").(types.AxolgoConfig)
	axolgoConfig.Logging.LogLevelVerbosity = 3
	viper.Set("axolgo-config", axolgoConfig)

	t.Run("init with environment variable", func(t *testing.T) {
		assert.NotPanics(t, func() { InitAxolgoConfig("") })
	})
}
