output "storage_account_id" {
  description = "ID of the created Storage Account"
  value       = azurerm_storage_account.storage.id
}

output "storage_account_primary_access_key" {
  description = "Primary access key of the Storage Account"
  value       = azurerm_storage_account.storage.primary_access_key
  sensitive   = true
}

output "storage_account_primary_connection_string" {
  description = "Primary connection string for the Storage Account"
  value       = azurerm_storage_account.storage.primary_connection_string
  sensitive   = true
}

output "container_id" {
  description = "ID of the Storage Container"
  value       = azurerm_storage_container.container.id
}
