variable "application_name" {} // used in Name tags on resources

variable "application_stage" {} // e.g test, dev, prod. Used in tags

variable "region" {} // AWS region to deploy resources to

variable "lambda_directory" { }

variable "temp_directory" {}