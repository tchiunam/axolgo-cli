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

	"k8s.io/klog/v2"

	"github.com/spf13/cobra"
	"github.com/tchiunam/axolgo-lib/cryptography"
)

var (
	genPassphraseLong = "Generate a passphrase."

	genPassphraseExample = `  # Generate a passphrase."
  axolgo crytography genPassphrase
`
)

// GenPassphraseOptions defines flags and other configuration parameters for the `genPassphrasencrypt` command
type GenPassphraseOptions struct {
	SaveFile string
}

// NewCmdGenPassphrase creates the `genPassphrase` command
func NewCmdGenPassphrase(ctx *context.Context) *cobra.Command {
	o := GenPassphraseOptions{}

	cmd := &cobra.Command{
		Use:                   "genPassphrase [-s]",
		DisableFlagsInUseLine: true,
		Short:                 "Generate a passphrase.",
		Long:                  genPassphraseLong,
		Example:               genPassphraseExample,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.complete(ctx, cmd, args); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&o.SaveFile, "save-file", "s", "", "Save to a file.")

	cmd.MarkFlagRequired("save-file")

	return cmd

}

// Complete takes the command arguments and execute.
func (o *GenPassphraseOptions) complete(_ *context.Context, _ *cobra.Command, args []string) error {
	if passphrase, err := cryptography.GeneratePassphrase(50); err == nil {
		if err = os.WriteFile(o.SaveFile, []byte(passphrase), 0644); err != nil {
			klog.Errorf("Failed to write passphrase into file: %s", o.SaveFile)
			return err
		}
	} else {
		klog.Errorf("Failed to generate passphrase: %s", err)
		return err
	}

	return nil
}
