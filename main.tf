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

module "aws-subnet" {
  source = "./aws-subnet"

  subnet_cidr_block = var.subnet_cidr_block
  vpc_id            = var.vpc_id
  subnet_name       = var.subnet_name
  aws_region        = var.aws_region
  aws_role_arn      = var.aws_role_arn
  subnet_tags       = var.subnet_tags
}
