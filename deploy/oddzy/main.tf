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


// Get latest AMI from repository???


### RESOURCES ###


// EC2 Instance


// VPC


// SUBNET - Public


// SUBNET - Private


// Internet gateway


// Route 53 zone


// Route 53 record


// ECS Cluster


// ECS Service


// ECS Task Definition (1 per container)


// ECS Container (1 per container)

