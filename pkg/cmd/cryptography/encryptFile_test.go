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
	"github.com/tchiunam/axolgo-cli/pkg/util"
)

// Initialize environment for the tests
func init() {
	util.InitAxolgoConfig(filepath.Join(filepath.Dir(""), "..", "..", "testdata", "config"))
}

// Clean up the test files
func _cleanTestEncryptFile(
	encOutputFilename string) {
	// Delete the encrypted file if it exists
	if _, err := os.Stat(encOutputFilename); err == nil {
		os.Remove(encOutputFilename)
	}
}

// TestNewCmdEncryptFile tests the NewCmdEncryptFile function
// to make sure it returns a valid command.
func TestNewCmdEncryptFile(t *testing.T) {
	cases := map[string]struct {
		use            string
		short          string
		hasFlags       bool
		keyFile        string
		filePath       string
		outputFilePath string
	}{
		"valid command": {
			use:            "encryptFile [-k] -f FILENAME -o OUTPUT_FILENAME",
			short:          "Encrypt a file.",
			hasFlags:       true,
			keyFile:        filepath.Join("testdata", "secret-test.key"),
			filePath:       filepath.Join("testdata", "story.txt"),
			outputFilePath: "",
		},
		"valid command with output file": {
			use:            "encryptFile [-k] -f FILENAME -o OUTPUT_FILENAME",
			short:          "Encrypt a file.",
			hasFlags:       true,
			keyFile:        filepath.Join("testdata", "secret-test.key"),
			filePath:       filepath.Join("testdata", "story.txt"),
			outputFilePath: filepath.Join("testdata", "story-encrypted.txt"),
		},
	}

	for name, c := range cases {
		_cleanTestEncryptFile(filepath.Join("testdata", "story-encrypted.txt"))
		t.Run(name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = []string{
				"cmd",
				"--key-file", c.keyFile,
				"--file", c.filePath}
			if c.outputFilePath != "" {
				os.Args = append(os.Args, "--output-file", c.outputFilePath)
			}

			cmd := NewCmdEncryptFile(nil)
			assert.Equal(t, c.use, cmd.Use)
			assert.Equal(t, c.short, cmd.Short)
			assert.Equal(t, c.hasFlags, cmd.Flags().HasFlags())
			assert.NoError(t, cmd.Execute())
		})
		_cleanTestEncryptFile(filepath.Join("testdata", "story-encrypted.txt"))
	}
}

// TestNewCmdEncryptFileInvalid calls the NewCmdEncryptFile function with an invalid
// input and verifies that an error is returned.
func TestNewCmdEncryptFileInvalid(t *testing.T) {
	cases := map[string]struct {
		use      string
		short    string
		hasFlags bool
		keyFile  string
		filePath string
	}{
		"non-exist key file": {
			keyFile:  filepath.Join("testdata", "missing.key"),
			filePath: filepath.Join("testdata", "story.txt"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = []string{
				"cmd",
				"--key-file", c.keyFile,
				"--file", c.filePath}

			cmd := NewCmdEncryptFile(nil)
			assert.Panics(t, func() { cmd.Execute() })
		})
	}
}
