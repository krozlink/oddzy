# Latest Amazon ECS optimized AMI
data "aws_ami" "ecs_ami" {
  filter {
    name   = "state"
    values = ["available"]
  }

  filter {
    name   = "name"
    values = ["amzn-ami-*-amazon-ecs-optimized"]
  }

  most_recent = true
}

// ECS Cluster
resource "aws_ecs_cluster" "main" {
  name = "${var.application_name}-${var.application_stage}"
}

// Launch Configuration
resource "aws_launch_configuration" "ecs_launch_configuration" {
  image_id             = "${data.aws_ami.ecs_ami.id}"
  instance_type        = "t2.medium"
  iam_instance_profile = "${aws_iam_instance_profile.ecs_instance_profile.id}"

  root_block_device {
    volume_type           = "standard"
    volume_size           = "20"
    delete_on_termination = true
  }

  lifecycle {
    create_before_destroy = true
  }

  security_groups             = ["${aws_security_group.main_sg.id}"]
  associate_public_ip_address = false
  key_name                    = "${var.ec2_key_pair}"

  user_data = <<EOF
#cloud-config
repo_update: true
repo_upgrade: all

packages:
  - amazon-efs-utils

runcmd:
  - echo ECS_CLUSTER=${aws_ecs_cluster.main.name} >> /etc/ecs/ecs.config
  - file_system_id_01=fs-1da46724
  - efs_directory=/mnt/efs
  - mkdir -p $${efs_directory}
  - echo "$${file_system_id_01}:/ $${efs_directory} efs tls,_netdev" >> /etc/fstab
  - mount -a -t efs defaults
  - chmod 777 $${efs_directory}
EOF
}

// Auto-Scaling Group
resource "aws_autoscaling_group" "ecs_autoscaling_group" {
  max_size             = "${var.ecs_max_instance_size}"
  min_size             = "${var.ecs_min_instance_size}"
  desired_capacity     = "${var.ecs_desired_capacity}"
  vpc_zone_identifier  = ["${aws_subnet.private.id}"]
  launch_configuration = "${aws_launch_configuration.ecs_launch_configuration.name}"
  target_group_arns    = ["${aws_lb_target_group.public.arn}", "${aws_lb_target_group.private.arn}"]
  health_check_type    = "ELB"
}

// IAM

resource "aws_iam_role" "ecs_service_role" {
  path               = "/"
  assume_role_policy = "${data.aws_iam_policy_document.ecs_service_policy.json}"
}

resource "aws_iam_role_policy_attachment" "ecs_service_role_attachment" {
  role       = "${aws_iam_role.ecs_service_role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
}

data "aws_iam_policy_document" "ecs_service_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ecs_instance_role" {
  path               = "/"
  assume_role_policy = "${data.aws_iam_policy_document.ecs_instance_policy.json}"
}

data "aws_iam_policy_document" "ecs_instance_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "ecs_instance_role_attachment" {
  role       = "${aws_iam_role.ecs_instance_role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_instance_profile" "ecs_instance_profile" {
  path = "/"
  role = "${aws_iam_role.ecs_instance_role.id}"
}
