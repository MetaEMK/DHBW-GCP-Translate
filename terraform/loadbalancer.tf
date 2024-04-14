resource "google_compute_health_check" "default" {
  name               = "http-basic-check"
  timeout_sec        = 5
  check_interval_sec = 10
  http_health_check {
    port = 80
    request_path = "/"
  }
}

resource "google_compute_backend_service" "default" {
  name        = "backend-service"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_health_check.default.self_link]

  backend {
    group = google_compute_instance_group.manager.self_link
  }
}

resource "google_compute_instance_group" "manager" {
  name        = "translator-group"
  zone        = local.ZONE
  #network     = "default"

  named_port {
    name = "http"
    port = 80
  }

  instances = [for instance in google_compute_instance.translator : instance.self_link]
}

resource "google_compute_url_map" "default" {
  name            = "web-map"
  default_service = google_compute_backend_service.default.self_link
}

resource "google_compute_target_http_proxy" "default" {
  name    = "http-lb-proxy"
  url_map = google_compute_url_map.default.self_link
}

resource "google_compute_global_forwarding_rule" "default" {
  name       = "http-content-rule"
  target     = google_compute_target_http_proxy.default.self_link
  port_range = "80"
}

resource "google_compute_firewall" "allow_health_check" {
  name    = "allow-health-check"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  source_ranges = ["130.211.0.0/22", "35.191.0.0/16"]  # Google Health Check IPs
  target_tags   = ["http-server"]  # Stellen Sie sicher, dass Ihre VMs diesen Tag haben
}

