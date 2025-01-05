provider "google" {
  project = var.gcp_project_id
  region  = var.gcp_region
}

resource "google_container_cluster" "primary" {
  name               = "go-test-cluster"
  location           = var.gcp_region
  initial_node_count = 1

  deletion_protection = false # This enables deletion protection for the cluster
  
  node_config {
    machine_type = "e2-medium"
  }
}


