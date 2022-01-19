terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "honestbank"
    workspaces {
      prefix = "subnet-"
    }
  }
}
