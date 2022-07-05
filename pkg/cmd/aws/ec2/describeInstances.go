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
	"bytes"
	"context"

	"k8s.io/klog/v2"

	"text/template"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	awsec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tchiunam/axolgo-cli/pkg/types"
	"github.com/tchiunam/axolgo-cloud/aws/ec2"
	"github.com/tchiunam/axolgo-cloud/aws/util"
	axolgolibutil "github.com/tchiunam/axolgo-lib/util"
)

var (
	describeInstancesLong = `Describe EC2 instances which are filterd by
given critera.
`
	describeInstancesExample = `  # Describe an EC2 instance
  axolgo aws ec2 describeInstances --instance-id i-831ao9b7co029d3ef
`
)

// DescribeInstancesOptions defines flags and other configuration parameters for the `describeInstances` command
type DescribeInstancesOptions struct {
	InstanceIDs            []string
	PrivateIPAddresses     []string
	PublicIPAddresses      []string
	SecurityGroupIDs       []string
	IamInstanceProfileArns []string
	MaxResults             int32
}

// NewCmdDescribeInstances creates the `describeInstances` command
func NewCmdDescribeInstances(ctx *context.Context) *cobra.Command {
	o := DescribeInstancesOptions{}

	cmd := &cobra.Command{
		Use:                   "describeInstances [-i] [-a] [-b] [-s] [-m] [-r]",
		DisableFlagsInUseLine: true,
		Short:                 "Describe EC2 instances.",
		Long:                  describeInstancesLong,
		Example:               describeInstancesExample,
		Run: func(cmd *cobra.Command, args []string) {
			if err := o.complete(ctx, cmd, args); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringArrayVarP(&o.InstanceIDs, "instance-id", "i", nil, "Instance IDs.")
	cmd.Flags().StringArrayVarP(&o.PrivateIPAddresses, "private-ip-address", "a", nil, "Private IP address.")
	cmd.Flags().StringArrayVarP(&o.PublicIPAddresses, "public-ip-address", "b", nil, "Public IP address.")
	cmd.Flags().StringArrayVarP(&o.SecurityGroupIDs, "security-group-id", "s", nil, "Security Group ID.")
	cmd.Flags().StringArrayVarP(&o.IamInstanceProfileArns, "iam-instance-profile.arn", "m", nil, "IAM instance profile's ARN.")
	cmd.Flags().Int32VarP(&o.MaxResults, "max-results", "r", 0, "Max. no. of records per batch.")

	return cmd

}

// Complete takes the command arguments and execute.
func (o *DescribeInstancesOptions) complete(ctx *context.Context, _ *cobra.Command, args []string) error {
	filterNVs := map[string][]string{
		"instance-id":              o.InstanceIDs,
		"private-ip-address":       o.PrivateIPAddresses,
		"ip-address":               o.PublicIPAddresses,
		"instance.group-id":        o.SecurityGroupIDs,
		"iam-instance-profile.arn": o.IamInstanceProfileArns,
	}

	// Built the Filter object as an input of AWS call
	var filters []awsec2types.Filter
	for filterName, filterValues := range filterNVs {
		klog.V(6).InfoS("create filter", "filterName", filterName, "filterValues", filterValues, "len(filterValues)", len(filterValues))
		if len(filterValues) != 0 {
			filters = append(filters,
				awsec2types.Filter{
					Name:   aws.String(filterName),
					Values: filterValues,
				})
		}
	}
	// MaxResults has a weird behvior. When it is specified,
	// The AWS client may not return any records even if there
	// is only one match. This happened to the case when
	// only Private IP Address or Security Group ID is given as
	// filter.
	// A workaround is to not set MaxResults unless it's provided
	// in the flag.
	klog.V(6).InfoS("max returned results", "MaxResults", o.MaxResults)
	input := &awsec2.DescribeInstancesInput{Filters: filters}
	if o.MaxResults > 0 {
		input.MaxResults = &o.MaxResults
	}

	axolgoConfig := viper.Get("axolgo-config").(types.AxolgoConfig)
	_, result, err := ec2.RunDescribeInstances(input, util.WithRegion(axolgoConfig.AWS.Region))
	if err != nil {
		klog.Fatalf("Failed to describe instances: %v", err)
	}

	var outputStringBytesBuffer bytes.Buffer
	securityGroupIDTmpl, err := template.New("security").Parse(ec2.SecurityGroupIDTemplateString)
	if err != nil {
		klog.Fatalf("Encountered error when creating Security Group ID template string: %v", err)
	}
	tagsTmpl, err := template.New("tags").Parse(ec2.TagsTemplateString)
	if err != nil {
		klog.Fatalf("Encountered error when creating Tag template string: %v", err)
	}

	klog.V(2).InfoS("result", "NextToken", result.NextToken, "len(Reservations)", len(result.Reservations))
	// Not handling NextToken. Do this when there is a need.
	for _, r := range result.Reservations {
		klog.Infof("Reservation ID: %v", *r.ReservationId)
		for _, i := range r.Instances {
			klog.Infof("    Instance ID: %v", *i.InstanceId)
			klog.Infof("    Private IP address: %v", *axolgolibutil.HushedStringPtr(i.PrivateIpAddress))
			klog.Infof("    Public IP address: %v", *axolgolibutil.HushedStringPtr(i.PublicIpAddress))
			if err := securityGroupIDTmpl.Execute(&outputStringBytesBuffer, i); err == nil {
				klog.Infof("    Security group IDs: [%v]", outputStringBytesBuffer.String())
				outputStringBytesBuffer.Reset()
			} else {
				klog.Errorf("Failed to extract Security Group IDs: %v", err)
			}
			if i.IamInstanceProfile == nil {
				klog.Info("    IAM instance profile ARN: []")
			} else {
				klog.Infof("    IAM instance profile ARN: [%v]", *axolgolibutil.HushedStringPtr(i.IamInstanceProfile.Arn))
			}
			klog.Infof("    Instance type: %v", i.InstanceType)
			if err := tagsTmpl.Execute(&outputStringBytesBuffer, i); err == nil {
				klog.Infof("    Tags: [%v]", outputStringBytesBuffer.String())
				outputStringBytesBuffer.Reset()
			} else {
				klog.Errorf("Failed to extract Tags: %v", err)
			}
		}
	}

	return err
}
