// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm

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
	registry.AddResourceTypeFactory("awscc_ssm_association", associationResourceType)
}

// associationResourceType returns the Terraform awscc_ssm_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSM::Association resource type.
func associationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"apply_only_at_cron_interval": {
			// Property: ApplyOnlyAtCronInterval
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"association_id": {
			// Property: AssociationId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique identifier of the association.",
			//   "examples": [
			//     "88df7b09-95e8-48c4-a3cb-08c2c20d5110",
			//     "203dd0ec-0055-4bf0-a872-707f72ef06aa"
			//   ],
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Unique identifier of the association.",
			Type:        types.StringType,
			Computed:    true,
		},
		"association_name": {
			// Property: AssociationName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the association.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the association.",
			Type:        types.StringType,
			Optional:    true,
		},
		"automation_target_parameter_name": {
			// Property: AutomationTargetParameterName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 50),
			},
		},
		"calendar_names": {
			// Property: CalendarNames
			// CloudFormation resource type schema:
			// {
			//   "examples": [
			//     [
			//       "calendar1",
			//       "calendar2"
			//     ],
			//     [
			//       "calendar3"
			//     ]
			//   ],
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"compliance_severity": {
			// Property: ComplianceSeverity
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "CRITICAL",
			//     "HIGH",
			//     "MEDIUM",
			//     "LOW",
			//     "UNSPECIFIED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"CRITICAL",
					"HIGH",
					"MEDIUM",
					"LOW",
					"UNSPECIFIED",
				}),
			},
		},
		"document_version": {
			// Property: DocumentVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of the SSM document to associate with the target.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The version of the SSM document to associate with the target.",
			Type:        types.StringType,
			Optional:    true,
		},
		"instance_id": {
			// Property: InstanceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the instance that the SSM document is associated with.",
			//   "examples": [
			//     "i-0e60836d21cf313c4",
			//     "mi-0532c22e49636ee13"
			//   ],
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the instance that the SSM document is associated with.",
			Type:        types.StringType,
			Optional:    true,
		},
		"max_concurrency": {
			// Property: MaxConcurrency
			// CloudFormation resource type schema:
			// {
			//   "examples": [
			//     "1%",
			//     "10%",
			//     "50%",
			//     "1"
			//   ],
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"max_errors": {
			// Property: MaxErrors
			// CloudFormation resource type schema:
			// {
			//   "examples": [
			//     "1%",
			//     "10%",
			//     "50%",
			//     "1"
			//   ],
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the SSM document.",
			//   "examples": [
			//     "AWS-GatherSoftwareInventory",
			//     "MyCustomSSMDocument"
			//   ],
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the SSM document.",
			Type:        types.StringType,
			Required:    true,
		},
		"output_location": {
			// Property: OutputLocation
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "S3Location": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "OutputS3BucketName": {
			//           "maxLength": 63,
			//           "minLength": 3,
			//           "type": "string"
			//         },
			//         "OutputS3KeyPrefix": {
			//           "maxLength": 1024,
			//           "type": "string"
			//         },
			//         "OutputS3Region": {
			//           "maxLength": 20,
			//           "minLength": 3,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"s3_location": {
						// Property: S3Location
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"output_s3_bucket_name": {
									// Property: OutputS3BucketName
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(3, 63),
									},
								},
								"output_s3_key_prefix": {
									// Property: OutputS3KeyPrefix
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 1024),
									},
								},
								"output_s3_region": {
									// Property: OutputS3Region
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(3, 20),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"parameters": {
			// Property: Parameters
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Parameter values that the SSM document uses at runtime.",
			//   "patternProperties": {
			//     "": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Parameter values that the SSM document uses at runtime.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.ListType{ElemType: types.StringType}},
			Optional: true,
		},
		"schedule_expression": {
			// Property: ScheduleExpression
			// CloudFormation resource type schema:
			// {
			//   "description": "A Cron or Rate expression that specifies when the association is applied to the target.",
			//   "examples": [
			//     "cron(0 0 */1 * * ? *)",
			//     "cron(0 16 ? * TUE *)",
			//     "rate(30 minutes)",
			//     "rate(7 days)"
			//   ],
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A Cron or Rate expression that specifies when the association is applied to the target.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
		},
		"sync_compliance": {
			// Property: SyncCompliance
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "AUTO",
			//     "MANUAL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AUTO",
					"MANUAL",
				}),
			},
		},
		"targets": {
			// Property: Targets
			// CloudFormation resource type schema:
			// {
			//   "description": "The targets that the SSM document sends commands to.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Values": {
			//         "items": {
			//           "type": "string"
			//         },
			//         "maxItems": 50,
			//         "minItems": 0,
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Values"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 5,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The targets that the SSM document sends commands to.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"values": {
						// Property: Values
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(0, 50),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 5,
				},
			),
			Optional: true,
		},
		"wait_for_success_timeout_seconds": {
			// Property: WaitForSuccessTimeoutSeconds
			// CloudFormation resource type schema:
			// {
			//   "maximum": 172800,
			//   "minimum": 15,
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(15, 172800),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::SSM::Association resource associates an SSM document in AWS Systems Manager with EC2 instances that contain a configuration agent to process the document.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::Association").WithTerraformTypeName("awscc_ssm_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"apply_only_at_cron_interval":      "ApplyOnlyAtCronInterval",
		"association_id":                   "AssociationId",
		"association_name":                 "AssociationName",
		"automation_target_parameter_name": "AutomationTargetParameterName",
		"calendar_names":                   "CalendarNames",
		"compliance_severity":              "ComplianceSeverity",
		"document_version":                 "DocumentVersion",
		"instance_id":                      "InstanceId",
		"key":                              "Key",
		"max_concurrency":                  "MaxConcurrency",
		"max_errors":                       "MaxErrors",
		"name":                             "Name",
		"output_location":                  "OutputLocation",
		"output_s3_bucket_name":            "OutputS3BucketName",
		"output_s3_key_prefix":             "OutputS3KeyPrefix",
		"output_s3_region":                 "OutputS3Region",
		"parameters":                       "Parameters",
		"s3_location":                      "S3Location",
		"schedule_expression":              "ScheduleExpression",
		"sync_compliance":                  "SyncCompliance",
		"targets":                          "Targets",
		"values":                           "Values",
		"wait_for_success_timeout_seconds": "WaitForSuccessTimeoutSeconds",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ssm_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
