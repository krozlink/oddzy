provider "aws" {
  region = "${var.region}"
}

terraform {
  backend "s3" {
    bucket = "oddzy"
    key    = "terraform/state"
    region = "ap-southeast-2"
  }
}

### DATA ###


// ECS Service


// ECS Task Definition (1 per container)


// ECS Container (1 per container)


// Route 53 zone


// Route 53 record

