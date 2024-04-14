resource "google_compute_subnetwork" "translator" {
  name = "translator"
  ip_cidr_range = "10.1.1.0/24"
  region = local.REGION
  network = "default"
}
