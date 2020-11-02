terraform {
  required_providers {
    adops = {
      versions = ["0.1"]
      source = "github.com/scastria/adops"
    }
  }
}

provider "adops" {
  username = "WWWW"
  password = "XXXX"
  organization = "YYYY"
  project = "ZZZZ"
}

data "adops_pipeline" "MyPipeline" {
  id = 227
}

output "MyOutput" {
  value = data.adops_pipeline.MyPipeline.name
}
