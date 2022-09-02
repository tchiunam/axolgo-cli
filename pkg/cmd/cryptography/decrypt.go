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
	decryptLong = "Decrypt the provided string with a key."

	decryptExample = `  # Decrypt a message with a key file.
  axolgo crytography decrypt --key-file secret.key --message 382b7f6eff497363265885f017286cd7ae7a417526419e8779abe535059bb6fd8cf58f11ef7214
`
)

// DecryptOptions defines flags and other configuration parameters for the `decrypt` command
type DecryptOptions struct {
	KeyFile string
	Message string
}

// NewCmdDecrypt creates the `decrypt` command
func NewCmdDecrypt(ctx *context.Context) *cobra.Command {
	o := DecryptOptions{}

	cmd := &cobra.Command{
		Use:                   "decrypt [-k] [-m]",
		DisableFlagsInUseLine: true,
		Short:                 "Decrypt a message.",
		Long:                  decryptLong,
		Example:               decryptExample,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.complete(ctx, cmd, args); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&o.KeyFile, "key-file", "k", "", "Key file.")
	cmd.Flags().StringVarP(&o.Message, "message", "m", "", "Message to be decrypted.")

	return cmd

}

// Complete takes the command arguments and execute.
func (o *DecryptOptions) complete(_ *context.Context, _ *cobra.Command, args []string) error {
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

	if message, err := hex.DecodeString(o.Message); err == nil {
		if data, err := cryptography.Decrypt([]byte(message), string(passphrase)); err == nil {
			fmt.Println(string(data))
		} else {
			klog.Errorf("Failed to decrypt message: %s", o.Message)
			return err
		}
	} else {
		klog.Errorf("Failed to decode message: %s", o.Message)
		return err
	}

	return nil
}
