variable "instance_name" {
  description = "Multipass instance name"
  type        = string
  default     = "k8s-lab-n1"
}

variable "cpus" {
  description = "Number of vCPUs"
  type        = number
  default     = 3
}

variable "memory" {
  description = "RAM size (e.g. 6G)"
  type        = string
  default     = "6G"
}

variable "disk" {
  description = "Disk size (e.g. 30G)"
  type        = string
  default     = "30G"
}

variable "image" {
  description = "Ubuntu image alias for Multipass"
  type        = string
  default     = "26.04"
}
