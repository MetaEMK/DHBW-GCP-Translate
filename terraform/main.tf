terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}

locals {
  REGION = "europe-north1"
  ZONE   = "europe-north1-a"
}

provider "google" {
  project = "cc2-translator"
  region  = local.REGION
  zone    = local.ZONE
}

