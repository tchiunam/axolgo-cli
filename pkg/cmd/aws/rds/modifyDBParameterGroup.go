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

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tchiunam/axolgo-cli/pkg/types"
	"github.com/tchiunam/axolgo-cloud/aws/rds"
	"github.com/tchiunam/axolgo-cloud/aws/util"
	"github.com/tchiunam/axolgo-lib/io/ioutil"
	axolgolibtypes "github.com/tchiunam/axolgo-lib/types"
)

var (
	modifyDBParameterGroupLong = `Modify DB Parameter Group with the parameters provided.
There are two types of parameters: static and dynamic. Parameters are
read from a yaml file. Example:

==============================
static:
  param1: value1
  param2: value2
  ...
dynamic:
  param1: value1
  param2: value2
  ...
==============================
`
	modifyDBParameterGroupExample = `  # Modify by providing a parameter file
  axolgo aws rds modifyDBParameterGroup -n standard_group -f parameters.yaml
`
)

// ModifyDBParameterGroupOptions defines flags and other configuration parameters for the `modifyDBParameterGroup` command
type ModifyDBParameterGroupOptions struct {
	Name          string
	ParameterFile string
}

// NewCmdModifyDBParameterGroup creates the `modifyDBParameterGroup` command
func NewCmdModifyDBParameterGroup(ctx *context.Context) *cobra.Command {
	o := ModifyDBParameterGroupOptions{}

	cmd := &cobra.Command{
		Use:                   "modifyDBParameterGroup -n NAME -f FILENAME",
		DisableFlagsInUseLine: true,
		Short:                 "Modify DB Parameter Group.",
		Long:                  modifyDBParameterGroupLong,
		Example:               modifyDBParameterGroupExample,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.complete(ctx, cmd, args); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&o.Name, "name", "n", "", "DB Group Name.")
	cmd.Flags().StringVarP(&o.ParameterFile, "parameter-file", "f", "", "The file that contains parameters.")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("parameter-file")

	return cmd

}

// Complete takes the command arguments and execute
func (o *ModifyDBParameterGroupOptions) complete(ctx *context.Context, _ *cobra.Command, args []string) error {
	yamlFile, err := ioutil.ReadYamlFile(o.ParameterFile)
	if err != nil {
		return err
	}
	yamlParameters := yamlFile.(*viper.Viper)

	parameters := [][]axolgolibtypes.Parameter{make([]axolgolibtypes.Parameter, 0), make([]axolgolibtypes.Parameter, 0)}
	for i, t := range []string{"static", "dynamic"} {
		if section := yamlParameters.Get(t); section != nil {
			for k, v := range section.(map[string]interface{}) {
				parameters[i] = append(
					parameters[i],
					axolgolibtypes.Parameter{
						Name:  aws.String(k),
						Value: aws.String(fmt.Sprintf("%v", v)),
					})
			}
		}
	}

	axolgoConfig := viper.Get("axolgo-config").(types.AxolgoConfig)
	_, err = rds.RunModifyDBParameterGroup(o.Name, parameters[0], parameters[1], util.WithRegion(axolgoConfig.AWS.Region))

	return err
}
