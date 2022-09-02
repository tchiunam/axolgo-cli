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
	"bufio"
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
	"k8s.io/klog/v2"

	"github.com/spf13/cobra"
	"github.com/tchiunam/axolgo-lib/cryptography"
)

var (
	encryptLong = "Encrypt the provided string with a key."

	encryptExample = `  # Encrypt the string "Hello World" with a key file.
  axolgo crytography encrypt --key-file secret.key --message "Hello World"
`
)

// EncryptOptions defines flags and other configuration parameters for the `encrypt` command
type EncryptOptions struct {
	KeyFile string
	Message string
}

// NewCmdEncrypt creates the `encrypt` command
func NewCmdEncrypt(ctx *context.Context) *cobra.Command {
	o := EncryptOptions{}

	cmd := &cobra.Command{
		Use:                   "encrypt [-k] [-m]",
		DisableFlagsInUseLine: true,
		Short:                 "Encrypt a message.",
		Long:                  encryptLong,
		Example:               encryptExample,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.complete(ctx, cmd, args); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&o.KeyFile, "key-file", "k", "", "Key file.")
	cmd.Flags().StringVarP(&o.Message, "message", "m", "", "Message to be encrypted.")

	return cmd
}

// Complete takes the command arguments and execute.
func (o *EncryptOptions) complete(_ *context.Context, _ *cobra.Command, args []string) error {
	var passphrase []byte
	var err error
	if o.KeyFile == "" {
		fmt.Print("Enter passphrase: ")
		if passphrase, err = term.ReadPassword(int(syscall.Stdin)); err != nil {
			klog.Errorf("Failed to read passphrase from stdin: %v", err)
			return err
		}
		fmt.Println()
	} else {
		passphrase, err = os.ReadFile(o.KeyFile)
		if err != nil {
			klog.Errorf("Failed to read key file: %s", o.KeyFile)
			return err
		}
	}
	if o.Message == "" {
		fmt.Println("Enter message to be encrypted. Enter a new line and press Ctrl+D to finish:")
		scanner := bufio.NewScanner(os.Stdin)
		input := make([]string, 0)
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}
		o.Message = strings.Join(input, "\n")
	}
	data, err := cryptography.Encrypt([]byte(o.Message), string(passphrase))
	fmt.Println(hex.EncodeToString(data))

	return nil
}
