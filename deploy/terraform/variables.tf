variable "region" {} // AWS region to deploy resources to

variable "repository" {} // ECR repository to retrieve docker images

variable "application_name" {} // used in Name tags on resources

variable "application_stage" {} // e.g test, dev, prod. Used in tags

variable "ec2_key_pair" {} // KeyPair to use for logging into instances

variable "run_ec2_instance" {} // Set to false to deploy resources without the EC2 instance / ECS containers

variable "efs_volume" {} // EFS volume to be used by ECS containers for persistent data - e.g elasticsearch logs, grafana, databases

variable "domain_name" {} // Domain name to use. Must be a hosted zone with this domain name

variable "bucket_name" {} // Bucket used for any storage required, e.g logs

variable "test_instance" {} // Set to true to deploy a publicly-accessible jump box instance for testing

variable "run_core_tasks" {} // Set to true to deploy core ECS containers - e.g ELK stack, nginx, grafana / prometheus

variable "run_app_tasks" {} // Set to true to deploy app ECS containers
