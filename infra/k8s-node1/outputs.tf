output "instance_name" {
  description = "Multipass instance name"
  value       = multipass_instance.k8s_node1.name
}

output "ipv4" {
  description = "Instance IP address"
  value       = multipass_instance.k8s_node1.ipv4
}
