resource aws_lb "main" {
  internal           = false
  load_balancer_type = "application"
  subnets            = ["${aws_subnet.public-a.id}", "${aws_subnet.public-b.id}"]
  security_groups    = ["${aws_security_group.load_balancer.id}"]

  access_logs {
    bucket  = "${var.bucket_name}"
    prefix  = "elb_logs"
    enabled = true
  }

  tags {
    Name = "${var.application_name}"
  }
}

resource aws_lb_target_group "public" {
  port     = 80
  protocol = "HTTP"
  vpc_id   = "${aws_vpc.main.id}"

  tags {
    Name = "${var.application_name}-public"
  }
}

resource aws_lb_target_group "private" {
  port     = 8080
  protocol = "HTTP"
  vpc_id   = "${aws_vpc.main.id}"

  tags {
    Name = "${var.application_name}-private"
  }
}

resource aws_lb_listener "main" {
  load_balancer_arn = "${aws_lb.main.arn}"
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = "${aws_lb_target_group.public.arn}"
    type             = "forward"
  }
}

resource aws_lb_listener_rule "internal" {
  listener_arn = "${aws_lb_listener.main.arn}"
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = "${aws_lb_target_group.private.arn}"
  }

  condition {
    field  = "host-header"
    values = ["internal.oddzy.xyz"]
  }
}

resource aws_security_group "load_balancer" {
  description = "ELB Security Group"
  vpc_id      = "${aws_vpc.main.id}"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource aws_s3_bucket_policy elb {
  bucket = "${var.bucket_name}"

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "ELBLogsPolicy",
  "Statement": [
    {
      "Sid": "AllowELBLogs",
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::${var.bucket_name}/elb_logs/*",
      "Principal": {
        "AWS": [
          "783225319266"
        ]
      }
    } 
  ]
}
  POLICY
}
