variable "region" {}

variable "repository" {} // ECR repository to retrieve docker images

variable "application_name" {}

variable "application_stage" {} // e.g test, dev, prod

variable "ecs_key_pair" {} // KeyPair to use for logging into ECS instances

variable "ecs_min_instance_size" {}

variable "ecs_max_instance_size" {}

variable "ecs_desired_capacity" {}

variable "oddzy_ami_tag" {}

variable "efs_volume" {} // EFS volume to be used by ECS containers for persistent data - e.g elasticsearch logs, grafana, databases

variable "nat_eip_allocation_id" {} // TODO: remove the variable and generate this

variable "domain_name" {} // Domain name to use. Must be a hosted zone with this domain name

variable "bucket_name" {} // Bucket used for any storage required, e.g logs
