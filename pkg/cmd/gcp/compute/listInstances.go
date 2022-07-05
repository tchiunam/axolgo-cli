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

package compute

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/klog/v2"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tchiunam/axolgo-cli/pkg/types"
	"github.com/tchiunam/axolgo-cloud/gcp/compute"
	"github.com/tchiunam/axolgo-cloud/gcp/util"
	axolgolibutil "github.com/tchiunam/axolgo-lib/util"
	"google.golang.org/api/iterator"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

var (
	listInstancesLong = `List compute engine instances which are filterd by
given critera.
`
	listInstancesExample = `  # List comput engine instance
  axolgo gcp compute listInstances --project proj1 --zone asia-east1-a --id 7452065390813417482
`
)

// ListInstancesOptions defines flags and other configuration parameters for the `listInstances` command
type ListInstancesOptions struct {
	Project    string
	Zone       string
	IDs        []string
	Names      []string
	MaxResults int32
}

// NewCmdListInstances creates the `listInstances` command
func NewCmdListInstances(ctx *context.Context) *cobra.Command {
	o := ListInstancesOptions{}

	cmd := &cobra.Command{
		Use:                   "listInstances -p [-z] [-i] [-n] [-r]",
		DisableFlagsInUseLine: true,
		Short:                 "List compute engine instances.",
		Long:                  listInstancesLong,
		Example:               listInstancesExample,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.complete(ctx, cmd, args); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&o.Project, "project", "p", "", "Project ID.")
	cmd.Flags().StringVarP(&o.Zone, "zone", "z", "", "Zone.")
	cmd.Flags().StringArrayVarP(&o.IDs, "id", "i", nil, "Instance IDs.")
	cmd.Flags().StringArrayVarP(&o.Names, "name", "n", nil, "Instance Names.")
	cmd.Flags().Int32VarP(&o.MaxResults, "max-results", "r", 0, "Max. no. of records per batch.")

	cmd.MarkFlagRequired("project")

	return cmd

}

// Complete takes the command arguments and execute.
func (o *ListInstancesOptions) complete(ctx *context.Context, _ *cobra.Command, args []string) error {
	filterNVs := map[string][]string{
		"id":   o.IDs,
		"name": o.Names,
	}

	var filters []string
	for filterName, filterValues := range filterNVs {
		klog.V(6).InfoS("process filter", "filterName", filterName, "filterValues", filterValues, "len(filterValues)", len(filterValues))
		if len(filterValues) != 0 {
			filters = append(
				filters,
				strings.Join(axolgolibutil.Map4s(filterValues, func(s string) string {
					return fmt.Sprintf("(%v = %v)", filterName, s)
				}), " or "))
		}
	}
	f := strings.Join(filters, " or ")

	axolgoConfig := viper.Get("axolgo-config").(types.AxolgoConfig)
	klog.V(3).InfoS("axolgoConfig", "GCP.GoogleApplicationCredentials", axolgoConfig.GCP.GoogleApplicationCredentials)

	// Use zone configured in the config file if not specified
	zone := o.Zone
	if zone == "" {
		zone = axolgoConfig.GCP.Zone
	}
	req := &computepb.ListInstancesRequest{
		Project: o.Project,
		Zone:    zone,
		Filter:  &f,
	}

	c, it, err := compute.RunListInstances(req, util.WithCredentialsFile(axolgolibutil.ExpandPath(axolgoConfig.GCP.GoogleApplicationCredentials)))
	if err != nil {
		klog.Fatalf("Failed to list compute engine instances: %v", err)
	}
	defer c.Close()

	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			klog.Fatalf("Failed to extract instance: %v", err)
		}
		klog.Infof("Name: %v", instance.GetName())
		klog.Infof("    ID: %v", instance.GetId())
		klog.Infof("    Zone: %v", instance.GetZone())
		for _, nic := range instance.GetNetworkInterfaces() {
			klog.Infof("    NetworkInterface: %v", nic.GetName())
			for _, ip := range nic.GetAccessConfigs() {
				klog.Infof("        IP: %v", ip.GetNatIP())
			}
		}
		for _, label := range instance.GetLabels() {
			klog.Infof("    Label: %v", label)
		}
	}

	return err
}
