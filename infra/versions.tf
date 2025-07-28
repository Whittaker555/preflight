terraform {
  backend "s3" {
    bucket = "preflight-state-bucket"
    key = "preflight-api/terraform.tfstate"
    region = "eu-west-2"
  }
  required_version = ">= 1.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}
