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

resource aws_lb "main" {
  internal           = false
  load_balancer_type = "application"
  subnets            = ["${aws_subnet.public.id}", "${aws_subnet.public-b.id}"]

  tags {
    Name = "${var.application_name}"
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
