variable "region" {}

variable "repository" {} // ECR repository to retrieve docker images

variable "application_name" {}

variable "application_stage" {} // e.g test, dev, prod

variable "ec2_key_pair" {} // KeyPair to use for logging into instances

variable "run_ec2_instance" {}

variable "oddzy_ami_tag" {}

variable "efs_volume" {} // EFS volume to be used by ECS containers for persistent data - e.g elasticsearch logs, grafana, databases

variable "domain_name" {} // Domain name to use. Must be a hosted zone with this domain name

variable "bucket_name" {} // Bucket used for any storage required, e.g logs

variable "test_instance" {}

variable "run_core_tasks" {}

variable "run_app_tasks" {}
