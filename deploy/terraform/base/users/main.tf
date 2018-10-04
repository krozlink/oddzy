resource "aws_dynamodb_table" "users" {
  name = "${var.application_name}"
  hash_key = "user_id"
  write_capacity = 1
  read_capacity = 1
  attribute {
    name = "user_id"
    type = "S"
  }
}