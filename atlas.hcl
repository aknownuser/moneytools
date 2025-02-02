variable "db_url" {
  type = string
  default = "postgres://moneytool:moneytool@localhost:5432/moneytool?sslmode=disable"
}

variable "dialect" {
  type = string
  default = "postgres" // | mysql | sqlite | sqlserver
}

variable "path_to_models" {
  type = string
  default = "./src/models"
}

data "external_schema" "gorm" {
  program = [
  "go",
  "run",
  "-mod=mod",
  "ariga.io/atlas-provider-gorm",
  "load",
  "--path", var.path_to_models,
  "--dialect", var.dialect, 
  ]
}


env "local" {
  src = data.external_schema.gorm.url
  url = var.db_url
  dev = "docker://postgres/16"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}