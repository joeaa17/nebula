terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.40.0"
    }
  }
  required_version = ">= 1.0"
}

provider "aws" {
  # AMI: ami-0caef02b518350c8b
  region = "eu-central-1"
  alias  = "eu_central_1"
}

provider "aws" {
  # AMI: ami-02ea247e531eb3ce6
  region = "us-west-1"
  alias  = "us_west_1"
}

provider "aws" {
  # AMI: ami-08c40ec9ead489470
  region = "us-east-1"
  alias  = "us_east_1"
}

provider "aws" {
  # AMI: ami-04b3c23ec8efcc2d6
  region = "sa-east-1"
  alias  = "sa_east_1"
}

provider "aws" {
  # AMI: ami-07651f0c4c315a529
  region = "ap-southeast-1"
  alias  = "ap_southeast_1"
}

variable "ssh_key" {
  type = string
}

locals {
  default_tags = {
    managed_by = "terraform"
    project    = "nebula-hydra-dial-down"
  }
}

module "nebula_eu_central_1_offset_0" {
  source = "./node"

  ami          = "ami-0caef02b518350c8b"
  ssh_key      = var.ssh_key
  offset       = 0
  region       = "eu-central-1"
  default_tags = local.default_tags

  providers = {
    aws = aws.eu_central_1
  }
}

output "eu_central_1_offset_0_ip_address" {
  value = module.nebula_eu_central_1_offset_0.ip_address
}

module "nebula_us_west_1_offset_1" {
  source = "./node"

  ami          = "ami-02ea247e531eb3ce6"
  ssh_key      = var.ssh_key
  offset       = 15
  region       = "us-west-1"
  default_tags = local.default_tags

  providers = {
    aws = aws.us_west_1
  }
}

output "us_west_1_offset_15_ip_address" {
  value = module.nebula_us_west_1_offset_1.ip_address
}



