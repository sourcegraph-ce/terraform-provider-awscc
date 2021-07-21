// Code generated by generators/resource/main.go; DO NOT EDIT.

package xray

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_xray_sampling_rule", samplingRule)
}

// samplingRule returns the Terraform aws_xray_sampling_rule resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::XRay::SamplingRule resource type.
func samplingRule(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"rule_arn": {
			// Property: RuleARN
			// PrimaryIdentifier: true
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			     "$ref": "#/definitions/RuleARN",
			     "type": "string"
			   }
			*/
			Description: `The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.`,
			Type:        types.StringType,
			Computed:    true,
		},
		"rule_name": {
			// Property: RuleName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			     "maxLength": 32,
			     "minLength": 1,
			     "$ref": "#/definitions/RuleName",
			     "type": "string"
			   }
			*/
			Description: `The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.`,
			Type:        types.StringType,
			Optional:    true,
		},
		"sampling_rule": {
			// Property: SamplingRule
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "Attributes": {
			         "additionalProperties": false,
			         "$comment": "String to string map",
			         "description": "Matches attributes derived from the request.",
			         "patternProperties": {
			           ".{1,}": {
			             "type": "string"
			           }
			         },
			         "type": "object"
			       },
			       "FixedRate": {
			         "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
			         "type": "number"
			       },
			       "HTTPMethod": {
			         "description": "Matches the HTTP method from a request URL.",
			         "maxLength": 10,
			         "type": "string"
			       },
			       "Host": {
			         "description": "Matches the hostname from a request URL.",
			         "maxLength": 64,
			         "type": "string"
			       },
			       "Priority": {
			         "description": "The priority of the sampling rule.",
			         "type": "integer"
			       },
			       "ReservoirSize": {
			         "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
			         "type": "integer"
			       },
			       "ResourceARN": {
			         "description": "Matches the ARN of the AWS resource on which the service runs.",
			         "maxLength": 500,
			         "type": "string"
			       },
			       "RuleARN": {
			         "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			         "$ref": "#/definitions/RuleARN",
			         "type": "string"
			       },
			       "RuleName": {
			         "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			         "maxLength": 32,
			         "minLength": 1,
			         "$ref": "#/definitions/RuleName",
			         "type": "string"
			       },
			       "ServiceName": {
			         "description": "Matches the name that the service uses to identify itself in segments.",
			         "maxLength": 64,
			         "type": "string"
			       },
			       "ServiceType": {
			         "description": "Matches the origin that the service uses to identify its type in segments.",
			         "maxLength": 64,
			         "type": "string"
			       },
			       "URLPath": {
			         "description": "Matches the path from a request URL.",
			         "maxLength": 128,
			         "type": "string"
			       },
			       "Version": {
			         "description": "The version of the sampling rule format (1)",
			         "type": "integer"
			       }
			     },
			     "$ref": "#/definitions/SamplingRule",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"attributes": {
						// Property: Attributes
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "$comment": "String to string map",
						     "description": "Matches attributes derived from the request.",
						     "patternProperties": {
						       ".{1,}": {
						         "type": "string"
						       }
						     },
						     "type": "object"
						   }
						*/
						Description: `Matches attributes derived from the request.`,
						// Pattern: ".{1,}"
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"fixed_rate": {
						// Property: FixedRate
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
						     "type": "number"
						   }
						*/
						Description: `The percentage of matching requests to instrument, after the reservoir is exhausted.`,
						Type:        types.NumberType,
						Optional:    true,
					},
					"http_method": {
						// Property: HTTPMethod
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the HTTP method from a request URL.",
						     "maxLength": 10,
						     "type": "string"
						   }
						*/
						Description: `Matches the HTTP method from a request URL.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"host": {
						// Property: Host
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the hostname from a request URL.",
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Description: `Matches the hostname from a request URL.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"priority": {
						// Property: Priority
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The priority of the sampling rule.",
						     "type": "integer"
						   }
						*/
						Description: `The priority of the sampling rule.`,
						Type:        types.NumberType,
						Optional:    true,
					},
					"reservoir_size": {
						// Property: ReservoirSize
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
						     "type": "integer"
						   }
						*/
						Description: `A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.`,
						Type:        types.NumberType,
						Optional:    true,
					},
					"resource_arn": {
						// Property: ResourceARN
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the ARN of the AWS resource on which the service runs.",
						     "maxLength": 500,
						     "type": "string"
						   }
						*/
						Description: `Matches the ARN of the AWS resource on which the service runs.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"rule_arn": {
						// Property: RuleARN
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						     "$ref": "#/definitions/RuleARN",
						     "type": "string"
						   }
						*/
						Description: `The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"rule_name": {
						// Property: RuleName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						     "maxLength": 32,
						     "minLength": 1,
						     "$ref": "#/definitions/RuleName",
						     "type": "string"
						   }
						*/
						Description: `The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"service_name": {
						// Property: ServiceName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the name that the service uses to identify itself in segments.",
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Description: `Matches the name that the service uses to identify itself in segments.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"service_type": {
						// Property: ServiceType
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the origin that the service uses to identify its type in segments.",
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Description: `Matches the origin that the service uses to identify its type in segments.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"url_path": {
						// Property: URLPath
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the path from a request URL.",
						     "maxLength": 128,
						     "type": "string"
						   }
						*/
						Description: `Matches the path from a request URL.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"version": {
						// Property: Version
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The version of the sampling rule format (1)",
						     "type": "integer"
						   }
						*/
						Description: `The version of the sampling rule format (1)`,
						Type:        types.NumberType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"sampling_rule_record": {
			// Property: SamplingRuleRecord
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "CreatedAt": {
			         "description": "When the rule was created, in Unix time seconds.",
			         "type": "string"
			       },
			       "ModifiedAt": {
			         "description": "When the rule was modified, in Unix time seconds.",
			         "type": "string"
			       },
			       "SamplingRule": {
			         "additionalProperties": false,
			         "properties": {
			           "Attributes": {
			             "additionalProperties": false,
			             "$comment": "String to string map",
			             "description": "Matches attributes derived from the request.",
			             "patternProperties": {
			               ".{1,}": {
			                 "type": "string"
			               }
			             },
			             "type": "object"
			           },
			           "FixedRate": {
			             "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
			             "type": "number"
			           },
			           "HTTPMethod": {
			             "description": "Matches the HTTP method from a request URL.",
			             "maxLength": 10,
			             "type": "string"
			           },
			           "Host": {
			             "description": "Matches the hostname from a request URL.",
			             "maxLength": 64,
			             "type": "string"
			           },
			           "Priority": {
			             "description": "The priority of the sampling rule.",
			             "type": "integer"
			           },
			           "ReservoirSize": {
			             "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
			             "type": "integer"
			           },
			           "ResourceARN": {
			             "description": "Matches the ARN of the AWS resource on which the service runs.",
			             "maxLength": 500,
			             "type": "string"
			           },
			           "RuleARN": {
			             "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			             "$ref": "#/definitions/RuleARN",
			             "type": "string"
			           },
			           "RuleName": {
			             "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			             "maxLength": 32,
			             "minLength": 1,
			             "$ref": "#/definitions/RuleName",
			             "type": "string"
			           },
			           "ServiceName": {
			             "description": "Matches the name that the service uses to identify itself in segments.",
			             "maxLength": 64,
			             "type": "string"
			           },
			           "ServiceType": {
			             "description": "Matches the origin that the service uses to identify its type in segments.",
			             "maxLength": 64,
			             "type": "string"
			           },
			           "URLPath": {
			             "description": "Matches the path from a request URL.",
			             "maxLength": 128,
			             "type": "string"
			           },
			           "Version": {
			             "description": "The version of the sampling rule format (1)",
			             "type": "integer"
			           }
			         },
			         "$ref": "#/definitions/SamplingRule",
			         "type": "object"
			       }
			     },
			     "$ref": "#/definitions/SamplingRuleRecord",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"created_at": {
						// Property: CreatedAt
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "When the rule was created, in Unix time seconds.",
						     "type": "string"
						   }
						*/
						Description: `When the rule was created, in Unix time seconds.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"modified_at": {
						// Property: ModifiedAt
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "When the rule was modified, in Unix time seconds.",
						     "type": "string"
						   }
						*/
						Description: `When the rule was modified, in Unix time seconds.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"sampling_rule": {
						// Property: SamplingRule
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Attributes": {
						         "additionalProperties": false,
						         "$comment": "String to string map",
						         "description": "Matches attributes derived from the request.",
						         "patternProperties": {
						           ".{1,}": {
						             "type": "string"
						           }
						         },
						         "type": "object"
						       },
						       "FixedRate": {
						         "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
						         "type": "number"
						       },
						       "HTTPMethod": {
						         "description": "Matches the HTTP method from a request URL.",
						         "maxLength": 10,
						         "type": "string"
						       },
						       "Host": {
						         "description": "Matches the hostname from a request URL.",
						         "maxLength": 64,
						         "type": "string"
						       },
						       "Priority": {
						         "description": "The priority of the sampling rule.",
						         "type": "integer"
						       },
						       "ReservoirSize": {
						         "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
						         "type": "integer"
						       },
						       "ResourceARN": {
						         "description": "Matches the ARN of the AWS resource on which the service runs.",
						         "maxLength": 500,
						         "type": "string"
						       },
						       "RuleARN": {
						         "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						         "$ref": "#/definitions/RuleARN",
						         "type": "string"
						       },
						       "RuleName": {
						         "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						         "maxLength": 32,
						         "minLength": 1,
						         "$ref": "#/definitions/RuleName",
						         "type": "string"
						       },
						       "ServiceName": {
						         "description": "Matches the name that the service uses to identify itself in segments.",
						         "maxLength": 64,
						         "type": "string"
						       },
						       "ServiceType": {
						         "description": "Matches the origin that the service uses to identify its type in segments.",
						         "maxLength": 64,
						         "type": "string"
						       },
						       "URLPath": {
						         "description": "Matches the path from a request URL.",
						         "maxLength": 128,
						         "type": "string"
						       },
						       "Version": {
						         "description": "The version of the sampling rule format (1)",
						         "type": "integer"
						       }
						     },
						     "$ref": "#/definitions/SamplingRule",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"attributes": {
									// Property: Attributes
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "$comment": "String to string map",
									     "description": "Matches attributes derived from the request.",
									     "patternProperties": {
									       ".{1,}": {
									         "type": "string"
									       }
									     },
									     "type": "object"
									   }
									*/
									Description: `Matches attributes derived from the request.`,
									// Pattern: ".{1,}"
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"fixed_rate": {
									// Property: FixedRate
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
									     "type": "number"
									   }
									*/
									Description: `The percentage of matching requests to instrument, after the reservoir is exhausted.`,
									Type:        types.NumberType,
									Optional:    true,
								},
								"http_method": {
									// Property: HTTPMethod
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Matches the HTTP method from a request URL.",
									     "maxLength": 10,
									     "type": "string"
									   }
									*/
									Description: `Matches the HTTP method from a request URL.`,
									Type:        types.StringType,
									Optional:    true,
								},
								"host": {
									// Property: Host
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Matches the hostname from a request URL.",
									     "maxLength": 64,
									     "type": "string"
									   }
									*/
									Description: `Matches the hostname from a request URL.`,
									Type:        types.StringType,
									Optional:    true,
								},
								"priority": {
									// Property: Priority
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The priority of the sampling rule.",
									     "type": "integer"
									   }
									*/
									Description: `The priority of the sampling rule.`,
									Type:        types.NumberType,
									Optional:    true,
								},
								"reservoir_size": {
									// Property: ReservoirSize
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
									     "type": "integer"
									   }
									*/
									Description: `A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.`,
									Type:        types.NumberType,
									Optional:    true,
								},
								"resource_arn": {
									// Property: ResourceARN
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Matches the ARN of the AWS resource on which the service runs.",
									     "maxLength": 500,
									     "type": "string"
									   }
									*/
									Description: `Matches the ARN of the AWS resource on which the service runs.`,
									Type:        types.StringType,
									Optional:    true,
								},
								"rule_arn": {
									// Property: RuleARN
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
									     "$ref": "#/definitions/RuleARN",
									     "type": "string"
									   }
									*/
									Description: `The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.`,
									Type:        types.StringType,
									Optional:    true,
								},
								"rule_name": {
									// Property: RuleName
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
									     "maxLength": 32,
									     "minLength": 1,
									     "$ref": "#/definitions/RuleName",
									     "type": "string"
									   }
									*/
									Description: `The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.`,
									Type:        types.StringType,
									Optional:    true,
								},
								"service_name": {
									// Property: ServiceName
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Matches the name that the service uses to identify itself in segments.",
									     "maxLength": 64,
									     "type": "string"
									   }
									*/
									Description: `Matches the name that the service uses to identify itself in segments.`,
									Type:        types.StringType,
									Optional:    true,
								},
								"service_type": {
									// Property: ServiceType
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Matches the origin that the service uses to identify its type in segments.",
									     "maxLength": 64,
									     "type": "string"
									   }
									*/
									Description: `Matches the origin that the service uses to identify its type in segments.`,
									Type:        types.StringType,
									Optional:    true,
								},
								"url_path": {
									// Property: URLPath
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "Matches the path from a request URL.",
									     "maxLength": 128,
									     "type": "string"
									   }
									*/
									Description: `Matches the path from a request URL.`,
									Type:        types.StringType,
									Optional:    true,
								},
								"version": {
									// Property: Version
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The version of the sampling rule format (1)",
									     "type": "integer"
									   }
									*/
									Description: `The version of the sampling rule format (1)`,
									Type:        types.NumberType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"sampling_rule_update": {
			// Property: SamplingRuleUpdate
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "Attributes": {
			         "additionalProperties": false,
			         "$comment": "String to string map",
			         "description": "Matches attributes derived from the request.",
			         "patternProperties": {
			           ".{1,}": {
			             "type": "string"
			           }
			         },
			         "type": "object"
			       },
			       "FixedRate": {
			         "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
			         "type": "number"
			       },
			       "HTTPMethod": {
			         "description": "Matches the HTTP method from a request URL.",
			         "maxLength": 10,
			         "type": "string"
			       },
			       "Host": {
			         "description": "Matches the hostname from a request URL.",
			         "maxLength": 64,
			         "type": "string"
			       },
			       "Priority": {
			         "description": "The priority of the sampling rule.",
			         "type": "integer"
			       },
			       "ReservoirSize": {
			         "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
			         "type": "integer"
			       },
			       "ResourceARN": {
			         "description": "Matches the ARN of the AWS resource on which the service runs.",
			         "maxLength": 500,
			         "type": "string"
			       },
			       "RuleARN": {
			         "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			         "$ref": "#/definitions/RuleARN",
			         "type": "string"
			       },
			       "RuleName": {
			         "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			         "maxLength": 32,
			         "minLength": 1,
			         "$ref": "#/definitions/RuleName",
			         "type": "string"
			       },
			       "ServiceName": {
			         "description": "Matches the name that the service uses to identify itself in segments.",
			         "maxLength": 64,
			         "type": "string"
			       },
			       "ServiceType": {
			         "description": "Matches the origin that the service uses to identify its type in segments.",
			         "maxLength": 64,
			         "type": "string"
			       },
			       "URLPath": {
			         "description": "Matches the path from a request URL.",
			         "maxLength": 128,
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/SamplingRuleUpdate",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"attributes": {
						// Property: Attributes
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "$comment": "String to string map",
						     "description": "Matches attributes derived from the request.",
						     "patternProperties": {
						       ".{1,}": {
						         "type": "string"
						       }
						     },
						     "type": "object"
						   }
						*/
						Description: `Matches attributes derived from the request.`,
						// Pattern: ".{1,}"
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"fixed_rate": {
						// Property: FixedRate
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
						     "type": "number"
						   }
						*/
						Description: `The percentage of matching requests to instrument, after the reservoir is exhausted.`,
						Type:        types.NumberType,
						Optional:    true,
					},
					"http_method": {
						// Property: HTTPMethod
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the HTTP method from a request URL.",
						     "maxLength": 10,
						     "type": "string"
						   }
						*/
						Description: `Matches the HTTP method from a request URL.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"host": {
						// Property: Host
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the hostname from a request URL.",
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Description: `Matches the hostname from a request URL.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"priority": {
						// Property: Priority
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The priority of the sampling rule.",
						     "type": "integer"
						   }
						*/
						Description: `The priority of the sampling rule.`,
						Type:        types.NumberType,
						Optional:    true,
					},
					"reservoir_size": {
						// Property: ReservoirSize
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
						     "type": "integer"
						   }
						*/
						Description: `A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.`,
						Type:        types.NumberType,
						Optional:    true,
					},
					"resource_arn": {
						// Property: ResourceARN
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the ARN of the AWS resource on which the service runs.",
						     "maxLength": 500,
						     "type": "string"
						   }
						*/
						Description: `Matches the ARN of the AWS resource on which the service runs.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"rule_arn": {
						// Property: RuleARN
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						     "$ref": "#/definitions/RuleARN",
						     "type": "string"
						   }
						*/
						Description: `The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"rule_name": {
						// Property: RuleName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
						     "maxLength": 32,
						     "minLength": 1,
						     "$ref": "#/definitions/RuleName",
						     "type": "string"
						   }
						*/
						Description: `The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"service_name": {
						// Property: ServiceName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the name that the service uses to identify itself in segments.",
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Description: `Matches the name that the service uses to identify itself in segments.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"service_type": {
						// Property: ServiceType
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the origin that the service uses to identify its type in segments.",
						     "maxLength": 64,
						     "type": "string"
						   }
						*/
						Description: `Matches the origin that the service uses to identify its type in segments.`,
						Type:        types.StringType,
						Optional:    true,
					},
					"url_path": {
						// Property: URLPath
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "Matches the path from a request URL.",
						     "maxLength": 128,
						     "type": "string"
						   }
						*/
						Description: `Matches the path from a request URL.`,
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "type": "string"
			         },
			         "Value": {
			           "type": "string"
			         }
			       },
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "$ref": "#/definitions/Tags",
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	schema := schema.Schema{
		Description: `This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.`,
		Version:     1,
		Attributes:  attributes,
	}

	var features ResourceTypeFeatures

	features |= ResourceTypeHasUpdatableAttribute

	resourceType, err := NewResourceType(
		"AWS::XRay::SamplingRule", // CloudFormation type name
		"aws_xray_sampling_rule",  // Terraform type name
		schema,                    // Terraform schema
		"/properties/RuleARN",     // Primary identifier property path (JSON Pointer)
		[]string{},                // Write-only property paths (JSON Pointer)
		features,
	)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema for %s:\n\n%v", "aws_xray_sampling_rule", schema)

	return resourceType, nil
}
