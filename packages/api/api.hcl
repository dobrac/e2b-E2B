variable "gcp_zone" {
  type = string
  default = "us-central1-a"
}

variable "image_name" {
  type = string
  default = ""
}

variable "api_port_name" {
  type = string
  default = ""
}

variable "api_port_number" {
  type = number
  default = 0
}

variable "consul_token" {
  type = string
  default = ""
}

variable "nomad_token" {
  type = string
  default = ""
}

variable "nomad_address" {
  type = string
  default = ""
}

variable "logs_proxy_address" {
  type = string
  default = ""
}

variable "supabase_url" {
  type = string
  default = ""
}

variable "supabase_key" {
  type = string
  default = ""
}

variable "api_admin_key" {
  type = string
  default = ""
}

job "orchestration-api" {
  datacenters = [var.gcp_zone]

  priority = 90

  group "api-service" {
    network {
      port "api" {
        static = var.api_port_number
      }
    }

    service {
      name = "api"
      port = var.api_port_number

      check {
        type     = "http"
        name     = "health"
        path     = "/health"
        interval = "20s"
        timeout  = "5s"
        port     = var.api_port_number
      }
    }

    task "start" {
      driver = "docker"

      resources {
        memory = 256
        memory_max = 2048
        cpu    = 500
      }

      env {
        LOGS_PROXY_ADDRESS = var.logs_proxy_address
        NOMAD_ADDRESS      = var.nomad_address
        NOMAD_TOKEN        = var.nomad_token
        CONSUL_TOKEN       = var.consul_token
        SUPABASE_URL       = var.supabase_url
        SUPABASE_KEY       = var.supabase_key
        API_ADMIN_KEY      = var.api_admin_key
      }

      config {
        network_mode = "host"
        image        = var.image_name
        ports        = [var.api_port_name]
        args = [
          "--port", "${var.api_port_number}",
        ]
      }
    }
  }
}