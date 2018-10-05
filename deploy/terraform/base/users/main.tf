resource "aws_dynamodb_table" "users" {
  name = "${var.application_name}-${var.application_stage}-users"
  hash_key = "user_id"
  write_capacity = 1
  read_capacity = 1
  attribute {
    name = "user_id"
    type = "S"
  }
}


resource "aws_iam_role" "lambda" {
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

