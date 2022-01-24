terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "honestbank"
    workspaces {
      name = "aws-subnet-main"
    }
  }
}
