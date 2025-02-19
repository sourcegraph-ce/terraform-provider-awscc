// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package devopsguru

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_devopsguru_resource_collection", resourceCollectionDataSource)
}

// resourceCollectionDataSource returns the Terraform awscc_devopsguru_resource_collection data source.
// This Terraform data source corresponds to the CloudFormation AWS::DevOpsGuru::ResourceCollection resource.
func resourceCollectionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ResourceCollectionFilter
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
		//	  "properties": {
		//	    "CloudFormation": {
		//	      "additionalProperties": false,
		//	      "description": "CloudFormation resource for DevOps Guru to monitor",
		//	      "properties": {
		//	        "StackNames": {
		//	          "description": "An array of CloudFormation stack names.",
		//	          "items": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "^[a-zA-Z*]+[a-zA-Z0-9-]*$",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 1000,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Tags": {
		//	      "description": "Tagged resources for DevOps Guru to monitor",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Tagged resource for DevOps Guru to monitor",
		//	        "properties": {
		//	          "AppBoundaryKey": {
		//	            "description": "A Tag key for DevOps Guru app boundary.",
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "TagValues": {
		//	            "description": "Tag values of DevOps Guru app boundary.",
		//	            "items": {
		//	              "maxLength": 256,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "maxItems": 1000,
		//	            "minItems": 1,
		//	            "type": "array"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"resource_collection_filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CloudFormation
				"cloudformation": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: StackNames
						"stack_names": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "An array of CloudFormation stack names.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "CloudFormation resource for DevOps Guru to monitor",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AppBoundaryKey
							"app_boundary_key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A Tag key for DevOps Guru app boundary.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TagValues
							"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "Tag values of DevOps Guru app boundary.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Tagged resources for DevOps Guru to monitor",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceCollectionType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of ResourceCollection",
		//	  "enum": [
		//	    "AWS_CLOUD_FORMATION",
		//	    "AWS_TAGS"
		//	  ],
		//	  "type": "string"
		//	}
		"resource_collection_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of ResourceCollection",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DevOpsGuru::ResourceCollection",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DevOpsGuru::ResourceCollection").WithTerraformTypeName("awscc_devopsguru_resource_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_boundary_key":           "AppBoundaryKey",
		"cloudformation":             "CloudFormation",
		"resource_collection_filter": "ResourceCollectionFilter",
		"resource_collection_type":   "ResourceCollectionType",
		"stack_names":                "StackNames",
		"tag_values":                 "TagValues",
		"tags":                       "Tags",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
