terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}

variable "SSH_KEY_PATH" {
  type = string
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

