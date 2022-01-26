# Terraform AWS Subnet module

A simple module that creates subnets for you.

## How to use:

Make sure you have a vpc and know what region you are in. 

Apply below module:

```terraform
module "aws-subnet" {
  source = "./aws-subnet"

  subnet_cidr_block = var.subnet_cidr_block
  vpc_id            = aws_vpc.main.id
  subnet_name       = var.subnet_name
  aws_region        = var.aws_region
  aws_role_arn      = var.aws_role_arn
  subnet_tags       = var.subnet_tags
}
```
`aws_role_arn` is optional and can be set to null (or omitted)

## Some built in features

* We automatically add the tag `CreatedBy = "terraform-aws-subnet"` to all subnets. As of now, this can't be turned off.
