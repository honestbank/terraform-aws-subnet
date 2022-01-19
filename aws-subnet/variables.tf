variable "subnet_cidr_block" {
    description = "CIDR block for subnet"
    default = null
}

variable "subnet_name" {
    description = "Name for subnet"
    default = null
}

variable "vpc_id" {
    description = "VPC ID"
    default = null
}

variable "aws_region" {
    description = "AWS Region"
    default = null
}

variable "aws_role_arn" {
    description = "AWS Role ARN"
    default = null
}

variable "subnet_tags" {
    description = "Tags"
    type = map
    default = {
        CreatedBy = "terraform-aws-subnet"
    } 
}