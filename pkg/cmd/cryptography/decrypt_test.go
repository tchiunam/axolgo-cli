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

package cryptography

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCmdDecrypt tests the NewCmdDecrypt function
// to make sure it returns a valid command.
func TestNewCmdDecrypt(t *testing.T) {
	cases := map[string]struct {
		use      string
		short    string
		hasFlags bool
		keyFile  string
		message  string
	}{
		"valid command": {
			use:      "decrypt [-k] [-m]",
			short:    "Decrypt a message.",
			hasFlags: true,
			keyFile:  filepath.Join("testdata", "secret-test.key"),
			message:  "a793ca5fc2eeb8b2c30cee2dd758ecb6663f4c84dde3912237fcddb346eb5f49e19f26472913f16ff6f8d1b220cba8d628a00c3ac6a57b294c2b440aad2276f52a8e62128adb7909",
		},
	}

	os.Setenv("AXOLGO_CONFIG_PATH", filepath.Join(filepath.Dir(""), "..", "..", "testdata", "config"))

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = []string{
				"cmd",
				"--key-file", c.keyFile,
				"--message", c.message}

			cmd := NewCmdDecrypt(nil)
			assert.Equal(t, c.use, cmd.Use)
			assert.Equal(t, c.short, cmd.Short)
			assert.Equal(t, c.hasFlags, cmd.Flags().HasFlags())
			assert.NoError(t, cmd.Execute())
		})
	}
}

// TestNewCmdDecryptInvalid calls the NewCmdDecrypt function with an invalid
// input and verifies that an error is returned.
func TestNewCmdDecryptInvalid(t *testing.T) {
	cases := map[string]struct {
		use      string
		short    string
		hasFlags bool
		keyFile  string
		message  string
	}{
		"non-exist key file": {
			keyFile: filepath.Join("testdata", "missing.key"),
			message: "a793ca5fc2eeb8b2c30cee2dd758ecb6663f4c84dde3912237fcddb346eb5f49e19f26472913f16ff6f8d1b220cba8d628a00c3ac6a57b294c2b440aad2276f52a8e62128adb7909",
		},
		"bad message": {
			keyFile: filepath.Join("testdata", "secret-test.key"),
			message: "123455",
		},
	}

	os.Setenv("AXOLGO_CONFIG_PATH", filepath.Join(filepath.Dir(""), "..", "..", "testdata", "config"))

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = []string{
				"cmd",
				"--key-file", c.keyFile,
				"--message", c.message}

			cmd := NewCmdDecrypt(nil)
			assert.Panics(t, func() { cmd.Execute() })
		})
	}
}
