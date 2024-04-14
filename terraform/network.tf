resource "google_compute_subnetwork" "translator" {
  name          = "translator"
  ip_cidr_range = "10.1.1.0/24"
  region        = local.REGION
  network       = "default"
}

resource "google_compute_firewall" "allow_traffic" {
  name    = "allow-traffic-from-to"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443"]
  }

  source_ranges = ["10.166.0.0/24"]
  target_tags   = ["allow-traffic"]
}

