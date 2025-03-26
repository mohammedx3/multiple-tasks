# Backstage

This Backstage template provisions Azure infrastructure using **Terraform**. It creates an **Azure Resource Group**, **Storage Account**, and **Storage Container**, and pushes the Terraform code to a GitHub repository. The template is fully customizable through parameters provided in the Backstage UI.

---

## Features

- **Resource Group**: Customizable name, location, and tags.
- **Storage Account**: Configurable name, access tier, replication type, and tags.
- **Storage Container**: Customizable name and access type.
- **GitHub Integration**: Automatically creates and pushes the Terraform configuration to a new GitHub repository.
- **Catalog Registration**: Registers the repository in Backstage for seamless integration.

---

## Prerequisites

- **GITHUB_TOKEN**: Ensure the `GITHUB_TOKEN` environment variable is set in your environment. This token is used to authenticate with GitHub and push the repository.

---

## Parameters

### **Azure Resource Group**
| Parameter               | Type   | Description                              | Default               |
|-------------------------|--------|------------------------------------------|-----------------------|
| `resource_group_name`   | String | Name of the Azure Resource Group         | `sre-challenge-flaschenpost` |
| `resource_group_location` | Enum  | Azure region for resource provisioning   | `westeurope`          |
| **Options**             |        | `westeurope`, `eastus`, `westus`, `northeurope` | |
| `resource_group_tag`    | String | Tag for the resource group               | `department = SRE`    |

---

### **Azure Storage Account**
| Parameter                  | Type   | Description                              | Default               |
|----------------------------|--------|------------------------------------------|-----------------------|
| `storage_account_name`     | String | Name of the Azure Storage Account (max 24 characters) | `sretaskforflaschenpost` |
| `storage_account_resource_group` | String | Resource group for the storage account  | `sre-challenge-flaschenpost` |
| `storage_access_tier`      | Enum   | Access tier for the storage account      | `Hot`                 |
| **Options**                |        | `Hot`, `Cool`                            |                       |
| `storage_account_tier`     | Enum   | Storage account tier                     | `Standard`            |
| **Options**                |        | `Standard`, `Premium`                    |                       |
| `storage_replication_type` | Enum   | Replication type for the storage account | `LRS`                 |
| **Options**                |        | `LRS`, `GRS`                             |                       |
| `storage_account_tag`      | String | Tag for the storage account              | `department = SRE`    |

---

### **Azure Storage Container**
| Parameter               | Type   | Description                              | Default               |
|-------------------------|--------|------------------------------------------|-----------------------|
| `container_name`        | String | Name of the storage container            | `sre`                 |
| `container_access_type` | Enum   | Access type for the container            | `private`             |
| **Options**             |        | `private`, `blob`, `container`           |                       |

---

### **GitHub Repository**
| Parameter               | Type   | Description                              | Default               |
|-------------------------|--------|------------------------------------------|-----------------------|
| `repoUrl`               | RepoUrlPicker | GitHub repository location            | N/A                   |
| `commitMessage`         | String | Commit message for the repository        | `"Add Azure Storage Terraform configuration"` |

---

## How the Template Works

1. **Generate Terraform Files**:
   - The template fetches Terraform files and replaces the `terraform.tfvars` values with the parameters provided by the user in the Backstage UI.

2. **Push to GitHub**:
   - A new GitHub repository is created with the Terraform configuration.
   - The Terraform code is committed and pushed to the repository.

3. **Register in Backstage**:
   - The repository is registered in the Backstage catalog for easy tracking.

---

## Steps

1. **Fetch Terraform Files**: 
   - Fetches the Terraform template files and substitutes user-provided values into `terraform.tfvars`.

2. **Publish to GitHub**: 
   - Pushes the generated Terraform code to a new GitHub repository specified by the `repoUrl` parameter.

3. **Register in Backstage**:
   - Adds the new repository to the Backstage catalog.

---

## Outputs

- **Repository**: A link to the newly created GitHub repository containing the Terraform files.
- **Catalog**: A link to the registered entity in the Backstage catalog.

---

## Notes

- Ensure the **GITHUB_TOKEN** environment variable is present before running this template.
- The template uses the provided parameters to dynamically generate and push the Terraform configuration.
- The default values for parameters can be updated in the Backstage UI during submission.
