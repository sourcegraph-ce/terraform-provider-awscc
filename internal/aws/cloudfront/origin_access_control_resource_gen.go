// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cloudfront_origin_access_control", originAccessControlResourceType)
}

// originAccessControlResourceType returns the Terraform awscc_cloudfront_origin_access_control resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFront::OriginAccessControl resource type.
func originAccessControlResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"origin_access_control_config": {
			// Property: OriginAccessControlConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Description": {
			//       "type": "string"
			//     },
			//     "Name": {
			//       "type": "string"
			//     },
			//     "OriginAccessControlOriginType": {
			//       "pattern": "^(s3)$",
			//       "type": "string"
			//     },
			//     "SigningBehavior": {
			//       "pattern": "^(never|no-override|always)$",
			//       "type": "string"
			//     },
			//     "SigningProtocol": {
			//       "pattern": "^(sigv4)$",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Name",
			//     "SigningProtocol",
			//     "SigningBehavior",
			//     "OriginAccessControlOriginType"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"description": {
						// Property: Description
						Type:     types.StringType,
						Optional: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
					},
					"origin_access_control_origin_type": {
						// Property: OriginAccessControlOriginType
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^(s3)$"), ""),
						},
					},
					"signing_behavior": {
						// Property: SigningBehavior
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^(never|no-override|always)$"), ""),
						},
					},
					"signing_protocol": {
						// Property: SigningProtocol
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("^(sigv4)$"), ""),
						},
					},
				},
			),
			Required: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::CloudFront::OriginAccessControl",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::OriginAccessControl").WithTerraformTypeName("awscc_cloudfront_origin_access_control")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                       "Description",
		"id":                                "Id",
		"name":                              "Name",
		"origin_access_control_config":      "OriginAccessControlConfig",
		"origin_access_control_origin_type": "OriginAccessControlOriginType",
		"signing_behavior":                  "SigningBehavior",
		"signing_protocol":                  "SigningProtocol",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
