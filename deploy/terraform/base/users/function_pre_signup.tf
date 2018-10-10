resource "aws_lambda_function" "pre_signup" {
  filename ="${var.temp_directory}/users-pre-signup.zip"
  function_name = "${var.application_name}-${var.application_stage}-users-pre-signup"
  handler = "users-pre-signup"
  role ="${aws_iam_role.lambda.arn}" 
  description = "Pre-Signup trigger for Cognito users"
  memory_size = 128
  timeout = 1
  runtime = "go1.x"
  source_code_hash = "${base64sha256(file("${var.temp_directory}/users-pre-signup.zip"))}"

  tags {
    Name = "${var.application_name}"
  }
}