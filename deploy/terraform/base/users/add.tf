resource "aws_lambda_function" "add" {
  filename ="${var.temp_directory}/users-add.zip"
  function_name = "${var.application_name}-${var.application_stage}-users-add"
  handler = "users-add"
  role ="${aws_iam_role.lambda.arn}" 
  description = "Creates a new user"
  memory_size = 128
  timeout = 1
  runtime = "go1.x"
  source_code_hash = "${base64sha256(file("${var.temp_directory}/users-add.zip"))}"

  tags {
    Name = "${var.application_name}"
  }
}

resource "null_resource" "build_add" {
  provisioner "local-exec" {
    command = "make -C ${var.lambda_directory} build"
  }

  triggers {
      user_exe = "${md5(file("${var.temp_directory}/users-add.zip"))}"
  }
}
