// Cloudwatch log group
resource aws_cloudwatch_log_group main {
  name = "${var.application_name}-${var.application_stage}"
}
