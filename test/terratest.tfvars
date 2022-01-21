aws_region        = "ap-southeast-1"
vpc_id            = "vpc-070389f4d969a4c96" # put here a VPC that exists in the account above.
subnet_cidr_block = "10.3.5.0/24"
subnet_name = "lab-networking-subnet"
subnet_tags = {
    CreatedBy = "terraform-aws-subnet-lab"
}
