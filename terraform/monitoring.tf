resource "google_compute_instance" "monitoring" {
  provider = google
  name = "monitoring-instance"
  machine_type = "e2-micro"

  metadata = {
    ssh-keys: format("jan:%s", file("${var.SSH_KEY_PATH}/monitoring.pub"))
  }

  boot_disk {
    auto_delete = true
    device_name = "monitoring"

    initialize_params {
      image = "projects/debian-cloud/global/images/debian-12-bookworm-v20240312"
      size  = 10
      type  = "pd-balanced"
    }

    mode = "READ_WRITE"
  }

  can_ip_forward      = false
  deletion_protection = false
  enable_display      = false

  labels = {
    goog-ec-src = "vm_add-tf"
  }


  network_interface {
    access_config {
      network_tier = "PREMIUM"
    }

    queue_count = 0
    stack_type  = "IPV4_ONLY"
    subnetwork  = "projects/${local.PROJECT}/regions/${local.REGION}/subnetworks/default"
  }

  scheduling {
    automatic_restart   = true
    on_host_maintenance = "MIGRATE"
    preemptible         = false
    provisioning_model  = "STANDARD"
  }

   service_account {
    email  = "390222210948-compute@developer.gserviceaccount.com"
    scopes = ["https://www.googleapis.com/auth/cloud-platform"]
   }

  shielded_instance_config {
    enable_integrity_monitoring = true
    enable_secure_boot          = false
    enable_vtpm                 = true
  }

  tags = ["http-server", "https-server"]
}

output "monitoring_ip" {
  value = google_compute_instance.monitoring.network_interface[0].access_config[0].nat_ip
}
