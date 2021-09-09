// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iotwireless_destinations", destinationsDataSourceType)
}

// destinationsDataSourceType returns the Terraform awscc_iotwireless_destinations data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoTWireless::Destination resource type.
func destinationsDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	// Required for acceptance testing.
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        providertypes.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::IoTWireless::Destination",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.FromCloudFormationAndTerraform("AWS::IoTWireless::Destination", "awscc_iotwireless_destinations", schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotwireless_destinations", "schema", hclog.Fmt("%v", schema))

	return pluralDataSourceType, nil
}
