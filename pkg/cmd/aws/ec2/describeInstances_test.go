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

package ec2

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tchiunam/axolgo-cli/pkg/util"
)

// Initialize environment for the tests
func init() {
	util.InitAxolgoConfig(filepath.Join(filepath.Dir(""), "..", "..", "..", "testdata", "config"))
}

// TestNewCmdDescribeInstances tests the NewCmdDescribeInstances function
// to make sure it returns a valid command.
func TestNewCmdDescribeInstances(t *testing.T) {
	cases := map[string]struct {
		use        string
		short      string
		hasFlags   bool
		instanceId string
	}{
		"valid command": {
			use:        "describeInstances [-i] [-a] [-b] [-s] [-m] [-r]",
			short:      "Describe EC2 instances.",
			hasFlags:   true,
			instanceId: "i-1234567890abcdef0",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = []string{
				"cmd",
				"--instance-id", c.instanceId}

			cmd := NewCmdDescribeInstances(nil)
			assert.Equal(t, c.use, cmd.Use)
			assert.Equal(t, c.short, cmd.Short)
			assert.Equal(t, c.hasFlags, cmd.Flags().HasFlags())
			assert.NoError(t, cmd.Execute())
		})
	}
}
