terraform {
  required_providers {
    multipass = {
      source  = "todoroff/multipass"
      version = "~> 1.0"
    }
  }
}

provider "multipass" {
  # Generous timeout: image download + package upgrades + kubeadm init
  command_timeout = 900
}

resource "multipass_instance" "k8s_node1" {
  name   = var.instance_name
  cpus   = var.cpus
  memory = var.memory
  disk   = var.disk
  image  = var.image

  cloud_init = file("${path.module}/cloud-init/cloud-config.yaml")
}
