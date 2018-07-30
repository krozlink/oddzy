// Application Load Balancer used to route traffic to the auto scaling group instances
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

// Load Balancer target group for public traffic
resource aws_lb_target_group "public" {
  port     = 80
  protocol = "HTTP"
  vpc_id   = "${aws_vpc.main.id}"

  health_check {
    interval            = 30
    path                = "/ping"
    port                = "traffic-port"
    protocol            = "HTTP"
    timeout             = 5
    healthy_threshold   = 5
    unhealthy_threshold = 3
    matcher             = "200"
  }

  tags {
    Name = "${var.application_name}-public"
  }
}

// Load Balancer target group for private traffic
resource aws_lb_target_group "private" {
  port     = 8080
  protocol = "HTTP"
  vpc_id   = "${aws_vpc.main.id}"

  health_check {
    interval            = 30
    path                = "/ping"
    port                = "traffic-port"
    protocol            = "HTTP"
    timeout             = 5
    healthy_threshold   = 5
    unhealthy_threshold = 3
    matcher             = "200"
  }

  tags {
    Name = "${var.application_name}-private"
  }
}

// SSL Certificate to use for HTTPS
data aws_acm_certificate "https" {
  domain      = "*.${var.domain_name}"
  statuses    = ["ISSUED"]
  most_recent = true
}

resource aws_lb_listener_certificate https {
  listener_arn    = "${aws_lb_listener.https.arn}"
  certificate_arn = "${data.aws_acm_certificate.https.arn}"
}

resource aws_lb_listener "https" {
  load_balancer_arn = "${aws_lb.main.arn}"
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = "${data.aws_acm_certificate.https.arn}"

  default_action {
    target_group_arn = "${aws_lb_target_group.public.arn}"
    type             = "forward"
  }
}

resource aws_lb_listener "http" {
  load_balancer_arn = "${aws_lb.main.arn}"
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = "${aws_lb_target_group.public.arn}"
    type             = "forward"
  }
}

// Load Balancer listener rule to forward internal.oddzy.xyz traffic to the private target group
resource aws_lb_listener_rule "internal-http" {
  listener_arn = "${aws_lb_listener.http.arn}"
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

resource aws_lb_listener_rule "internal-https" {
  listener_arn = "${aws_lb_listener.https.arn}"
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

// S3 bucket policy allowing the load balancer to write access logs
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
