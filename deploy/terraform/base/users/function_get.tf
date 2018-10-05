resource "aws_lambda_function" "get" {
  filename ="${var.temp_directory}/users-get.zip"
  function_name = "${var.application_name}-${var.application_stage}-users-get"
  handler = "users-get"
  role ="${aws_iam_role.lambda.arn}" 
  description = "Retrieves a user"
  memory_size = 128
  timeout = 1
  runtime = "go1.x"
  source_code_hash = "${base64sha256(file("${var.temp_directory}/users-get.zip"))}"

  tags {
    Name = "${var.application_name}"
  }
  
  depends_on = ["null_resource.build_get"]
}

resource "null_resource" "build_get" {
  provisioner "local-exec" {
    command = "make -C ${var.lambda_directory} build"
  }

  triggers {
      user_exe = "${uuid()}"
  }
}
