---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_opensearchserverless_security_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Amazon OpenSearchServerless security policy resource
---

# awscc_opensearchserverless_security_policy (Resource)

Amazon OpenSearchServerless security policy resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy` (String) The JSON policy document that is the content for the policy

### Optional

- `description` (String) The description of the policy
- `name` (String) The name of the policy
- `type` (String) The possible types for the network policy

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_opensearchserverless_security_policy.example <resource ID>
```
