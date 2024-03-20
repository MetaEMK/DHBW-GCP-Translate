terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}

provider "google" {
  project = "cc2-translator"
  region  = "europe-west1"
  zone    = "europe-west1-a"
}

