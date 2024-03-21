variable "DB_USERNAME" {
  type = string
}

variable "DB_PASSWORD" {
  type = string
}

variable "DB_NAME" {
  type = string
}

# Datenblock, um die externe IP der Compute Engine-Instanz zu erhalten
data "google_compute_instance" "translator_instance" {
  name = "translator-instance"
}

# Google Cloud SQL-Instanz
resource "google_sql_database_instance" "postgres-database" {
  name             = "translator_database"
  database_version = "POSTGRES_15"

  settings {
    tier = "db-f1-micro"
    # ip_configuration {
    #   authorized_networks {
    #     name  = "Compute Engine-Instanz"
    #     value = data.google_compute_instance.translator_instance.network_interface[0].access_config[0].nat_ip
    #   }
    # }
  }
}

# PostgreSQL-Datenbank
resource "google_sql_database" "default" {
  name     = var.DB_NAME
  instance = google_sql_database_instance.postgres-database.name
}

# PostgreSQL-Benutzer
resource "google_sql_user" "default" {
  name     = var.DB_USERNAME
  instance = google_sql_database_instance.postgres-database.name
  password = var.DB_PASSWORD
}

