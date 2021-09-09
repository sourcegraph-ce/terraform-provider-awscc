// Code generated by generators/resource/main.go; DO NOT EDIT.

package devopsguru

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_devopsguru_notification_channel", notificationChannelResourceType)
}

// notificationChannelResourceType returns the Terraform awscc_devopsguru_notification_channel resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DevOpsGuru::NotificationChannel resource type.
func notificationChannelResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"config": {
			// Property: Config
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about notification channels you have configured with DevOps Guru.",
			//   "properties": {
			//     "Sns": {
			//       "additionalProperties": false,
			//       "description": "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
			//       "properties": {
			//         "TopicArn": {
			//           "maxLength": 1024,
			//           "minLength": 36,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Information about notification channels you have configured with DevOps Guru.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sns": {
						// Property: Sns
						Description: "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"topic_arn": {
									// Property: TopicArn
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(36, 1024),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
			// Config is a force-new attribute.
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of a notification channel.",
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of a notification channel.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DevOpsGuru::NotificationChannel").WithTerraformTypeName("awscc_devopsguru_notification_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"config":    "Config",
		"id":        "Id",
		"sns":       "Sns",
		"topic_arn": "TopicArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_devopsguru_notification_channel", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
