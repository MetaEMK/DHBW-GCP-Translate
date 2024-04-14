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
  PROJECT = "cc2-translator"
}

provider "google" {
  project = local.PROJECT
  region  = local.REGION
  zone    = local.ZONE
}

variable "SSH_KEY_PATH" {
  type = string
}

