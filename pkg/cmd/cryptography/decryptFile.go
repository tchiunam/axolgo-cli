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
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
	"k8s.io/klog/v2"

	"github.com/spf13/cobra"
	"github.com/tchiunam/axolgo-lib/cryptography"
	"github.com/tchiunam/axolgo-lib/util"
)

var (
	decryptFileLong = "Decrypt the provided file with a key."

	decryptFileExample = `  # Decrypt the file 'example.txt with a key file.
  axolgo crytography decryptFile --key-file secret.key --file example.txt
`
)

// DecryptFileOptions defines flags and other configuration parameters for the `decryptFile` command
type DecryptFileOptions struct {
	KeyFile        string
	FilePath       string
	OutputFilePath string
}

// NewCmdDecryptFile creates the `decryptFile` command
func NewCmdDecryptFile(ctx *context.Context) *cobra.Command {
	o := DecryptFileOptions{}

	cmd := &cobra.Command{
		Use:                   "decryptFile [-k] -f FILENAME -o OUTPUT_FILENAME",
		DisableFlagsInUseLine: true,
		Short:                 "Decrypt a file.",
		Long:                  decryptFileLong,
		Example:               decryptFileExample,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.complete(ctx, cmd, args); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&o.KeyFile, "key-file", "k", "", "Key file.")
	cmd.Flags().StringVarP(&o.FilePath, "file", "f", "", "File to be decrypted.")
	cmd.Flags().StringVarP(&o.OutputFilePath, "output-file", "o", "", "Output file.")

	cmd.MarkFlagRequired("file")

	return cmd
}

// Complete takes the command arguments and execute.
func (o *DecryptFileOptions) complete(_ *context.Context, _ *cobra.Command, args []string) error {
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

	var outputFilePath string
	if o.OutputFilePath == "" {
		outputFilePath = util.AddSuffixToFileName(o.FilePath, "-decrypted")
	} else {
		outputFilePath = o.OutputFilePath
	}
	_, err = cryptography.DecryptFile(
		o.FilePath,
		string(passphrase),
		cryptography.WithOutputFilename(outputFilePath),
	)

	return err
}
