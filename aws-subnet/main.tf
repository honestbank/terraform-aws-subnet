provider "aws" {
  region = var.aws_region
  #skip_region_validation = var.aws_provider_skip_region_validation

  assume_role {
    role_arn = var.aws_role_arn
  }

  // This is necessary so that tags required for eks can be applied to the vpc without changes to the vpc wiping them out.
  // https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/resource-tagging
  ignore_tags {
    key_prefixes = ["kubernetes.io/"]
  }
}

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

data "aws_vpc" "vpc" {
  # use the vpc passed as input. If no input, use the one from remote state in network-infrastructure.
  id = var.vpc_id
}

resource "aws_subnet" "main" {
  vpc_id     = data.aws_vpc.vpc.id
  cidr_block = var.subnet_cidr_block

  map_public_ip_on_launch = false
  tags = merge({
    Name = var.subnet_name
    CreatedBy = "terraform-aws-subnet"
  }, var.subnet_tags)
}