provider "aws" {
  region = "${var.region}"
}

terraform {
  backend "s3" {
    bucket = "oddzy"
    key    = "terraform/state/base"
    region = "ap-southeast-2"
  }
}

// AWS Account ID
data "aws_caller_identity" "current" {}

module "users" {
    source = "./users"
    application_name = "${var.application_name}"
    application_stage = "${var.application_stage}"
    region = "${var.region}"
    lambda_directory = "${var.lambda_directory}"
    temp_directory = "${var.temp_directory}"
}