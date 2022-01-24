terraform {
  required_version = "~> 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    template = {
      source  = "hashicorp/template"
      version = "~> 2.2.0"
    }
  }
}


provider "aws" {
  region = var.aws_region
  alias  = "target"
  assume_role {
    role_arn = var.aws_role_arn
  }
}

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

module "aws-subnet" {
  source = "./aws-subnet"

  subnet_cidr_block = var.subnet_cidr_block
  vpc_id            = aws_vpc.main.id
  subnet_name       = var.subnet_name
  aws_region        = var.aws_region
  aws_role_arn      = var.aws_role_arn
  subnet_tags       = var.subnet_tags
}
