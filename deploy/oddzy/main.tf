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

data "aws_ami" "ecs_ami" {
  filter {
    name   = "state"
    values = ["available"]
  }

  filter {
    name   = "name"
    values = ["amzn-ami-*-amazon-ecs-optimized"]
  }

  most_recent = true
}

### RESOURCES ###

// EC2 Instance
resource "aws_instance" "server" {
  ami           = "${data.aws_ami.ecs_ami.id}"
  instance_type = "t2.medium"
  depends_on    = ["aws_internet_gateway.gw"]
}

// VPC
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

// SUBNET - Public
resource "aws_subnet" "subnet_public" {
  vpc_id     = "${aws_vpc.main.id}"
  cidr_block = "10.0.1.0/24"
}

// Internet gateway
resource "aws_internet_gateway" "gw" {
  vpc_id = "${aws_vpc.main.id}"
}

// Route 53 zone

// Route 53 record

// ECS Cluster
data "aws_ecs_cluster" "ecs_main" {
  cluster_name = "oddzy-dev-cluster"
}

// ECS Service


// ECS Task Definition (1 per container)


// ECS Container (1 per container)

