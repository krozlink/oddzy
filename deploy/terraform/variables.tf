variable "region" {
  type = "string"
}

variable "repository" {
  type = "string"
}

variable "application_name" {
  type = "string"
}

variable "application_stage" {
  type = "string"
}

variable "ecs_key_pair" {
  type = "string"
}

variable "ecs_min_instance_size" {
  type = "string"
}

variable "ecs_max_instance_size" {
  type = "string"
}

variable "ecs_desired_capacity" {
  type = "string"
}

variable "oddzy_ami_tag" {
  type = "string"
}

variable "efs_volume" {
  type = "string"
}

variable "nat_eip_allocation_id" {
  type = "string"
}

variable "hosted_zone" {
  type = "string"
}
