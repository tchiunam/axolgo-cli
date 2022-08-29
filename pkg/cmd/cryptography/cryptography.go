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
	"context"
	"os"

	"github.com/spf13/cobra"
)

// NewCryptographyCmd creates the `cryptography` command
func NewCryptographyCmd(ctx *context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cryptography",
		Short: "Cryptography utilities for securing the resources you manage",
		Long:  "There are many cryptography implementations we can choose. This is a set of utilities that picked the useful ones and is designed to help you focus on your business requirements.",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	cmd.AddCommand(
		NewCmdEncrypt(ctx),
		NewCmdDecrypt(ctx),
		NewCmdGenPassphrase(ctx),
	)

	return cmd
}

func init() {
}
