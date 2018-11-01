resource aws_kinesis_firehose_delivery_stream tracking {
    name= "${var.application_name}-${var.application_stage}-tracking"
    destination = "extended_s3"

    extended_s3_configuration {
        role_arn   = "${aws_iam_role.firehose_role.arn}"
        bucket_arn = "arn:aws:s3:::${var.bucket_name}"
        prefix="tracking-kinesis/"

        buffer_size = 5
        buffer_interval = 300

        cloudwatch_logging_options {
            enabled = true
            log_group_name = "${var.application_name}-${var.application_stage}"
            log_stream_name = "tracking"
        }
    }
}


resource "aws_iam_role" "firehose_role" {
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "firehose.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}