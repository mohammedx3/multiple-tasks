variable "resource_group_name" {
  description = "The name of the Azure Resource Group"
  type        = string
}

variable "resource_group_location" {
  description = "The Azure region where resources will be provisioned"
  type        = string
  default     = "westeurope"
}

variable "tags" {
  description = "Common tags for all resources"
  type        = map(string)
  default = {
    department = "SRE"
  }
}

variable "storage_account_name" {
  description = "The name of the Azure Storage Account"
  type        = string
}

variable "container_name" {
  description = "The name of the Storage Account Container"
  type        = string
}

variable "storage_access_tier" {
  description = "The access tier for the storage account"
  type        = string
  default     = "Hot"
}

variable "storage_account_tier" {
  description = "The tier of the storage account"
  type        = string
  default     = "Standard"
}

variable "storage_replication_type" {
  description = "The replication type for the storage account"
  type        = string
  default     = "LRS"
}

variable "container_access_type" {
  description = "The access type for the storage container"
  type        = string
  default     = "private"
}
