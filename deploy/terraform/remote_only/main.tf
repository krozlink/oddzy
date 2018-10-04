provider "aws" {
  region = "${var.region}"
}

terraform {
  backend "s3" {
    bucket = "oddzy"
    key    = "terraform/state/remote_only"
    region = "ap-southeast-2"
  }
}

// AWS Account ID
data "aws_caller_identity" "current" {}
