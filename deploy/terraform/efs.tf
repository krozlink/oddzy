resource aws_efs_mount_target container {
  file_system_id  = "${var.efs_volume}"
  subnet_id       = "${aws_subnet.private.id}"
  security_groups = ["${aws_security_group.efs.id}"]
}

resource aws_security_group efs {
  description = "Allow inbound for EFS"
  vpc_id      = "${aws_vpc.main.id}"

  ingress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    security_groups = ["${aws_security_group.main_sg.id}"]
  }
}
