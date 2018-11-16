resource aws_kinesis_firehose_delivery_stream tracking {
    name= "${var.application_name}-${var.application_stage}-tracking"
    destination = "extended_s3"

    extended_s3_configuration {
        role_arn   = "${aws_iam_role.firehose_role.arn}"
        bucket_arn = "arn:aws:s3:::${var.bucket_name}"
        prefix="tracking-kinesis/"

        buffer_size = 5
        buffer_interval = 60

        cloudwatch_logging_options {
            enabled = true
            log_group_name = "${aws_cloudwatch_log_group.main.name}"
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

resource "aws_iam_role_policy" "firehose_policy" {
  name = "test_policy"
  role = "${aws_iam_role.firehose_role.id}"

  policy = <<EOF
{
    "Version": "2012-10-17",  
    "Statement":
    [    
        {      
            "Effect": "Allow",      
            "Action": [        
                "s3:AbortMultipartUpload",        
                "s3:GetBucketLocation",        
                "s3:GetObject",        
                "s3:ListBucket",        
                "s3:ListBucketMultipartUploads",        
                "s3:PutObject"
            ],      
            "Resource": [        
                "arn:aws:s3:::oddzy",
                "arn:aws:s3:::oddzy/*"		    
            ]    
        },        
        {
            "Effect": "Allow",
            "Action": [
                "kinesis:DescribeStream",
                "kinesis:GetShardIterator",
                "kinesis:GetRecords"
            ],
            "Resource": "${aws_kinesis_firehose_delivery_stream.tracking.arn}"
        },
        {
           "Effect": "Allow",
           "Action": [
               "logs:PutLogEvents"
           ],
           "Resource": [
               "arn:aws:logs:${var.region}:${data.aws_caller_identity.current.account_id}:log-group:${aws_cloudwatch_log_group.main.name}:log-stream:tracking"
           ]
        }
    ]
}
EOF
}

resource aws_cloudwatch_log_stream tracking {
  name="tracking"
  log_group_name = "${aws_cloudwatch_log_group.main.name}"
}