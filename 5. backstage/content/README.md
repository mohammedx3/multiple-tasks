# Azure Resources

## Overview
This documentation outlines the Terraform configuration for provisioning Azure infrastructure resources. The setup includes the creation of a **Resource Group**, **Storage Account**, and **Storage Container** in the **Azure West Europe** region.

## Resources
### 1. **Azure Resource Group**
   - Acts as the central container for all Azure resources.

### 2. **Azure Storage Account**
   - Provides storage services to store objects/blobs.

### 3. **Azure Storage Container**
   - A container within the storage account to store blobs.

---

## Prerequisites
To use this Terraform configuration, ensure you have the following installed and configured:

1. **Terraform** (v1.0+ or [OpenTofu](https://opentofu.org/) v1.0+)
2. **Azure CLI** installed and authenticated with appropriate credentials.

---

## Configuration

### Required Variables
Create a `terraform.tfvars` file with the following variables:

```hcl
resource_group_name  = "sre-challenge-flaschenpost"
storage_account_name = "sretaskforflaschenpost"
container_name       = "sre"
```

> Note:
> Azure storage account names have a maximum limit of 24 characters. The name srechallengeforflaschenpost exceeds this limit, so it has been shortened to sretaskforflaschenpost.

### Optional Variables
You can override the following default variables in your `terraform.tfvars` file:

| Variable                  | Default Value           | Description                              |
|---------------------------|-------------------------|------------------------------------------|
| `resource_group_location` | `westeurope`           | The region for the resource group.       |
| `tags`                    | `{ department = "SRE" }` | Tags for resource organization.          |
| `storage_access_tier`     | `Hot`                  | The access tier for the storage account. |
| `storage_account_tier`    | `Standard`             | The performance tier of the storage.     |
| `storage_replication_type`| `LRS`                  | Replication type (Locally-Redundant).    |
| `container_access_type`   | `private`              | Access type for the storage container.   |

### Usage

```bash
terraform init
terraform plan -var-file=terraform.tfvars
terraform apply -var-file=terraform.tfvars
```


### Outputs
After provisioning, the following outputs will be available:

| Output Name                          | Description                                              | Sensitive |
|--------------------------------------|----------------------------------------------------------|-----------|
| `storage_account_id`                 | The ID of the Azure Storage Account                     | No        |
| `storage_account_primary_access_key` | The primary access key for the storage account          | Yes       |
| `storage_account_primary_connection_string` | The connection string for the storage account     | Yes       |
| `container_id`                       | The ID of the Azure Storage Container                   | No        |