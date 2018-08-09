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

// Launch Configuration for instances deployed in the auto scaling group
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

  depends_on = ["aws_efs_mount_target.container", "aws_s3_bucket_object.website"]

  user_data = <<EOF
#cloud-config
repo_update: true
repo_upgrade: all

packages:
  - amazon-efs-utils
  - aws-cli
  - unzip

runcmd:
  # Install SSM Agent
  - cd /tmp
  - yum install -y https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/linux_amd64/amazon-ssm-agent.rpm
  - start amazon-ssm-agent

  # Increase max number of memory map areas - this is required by elasticsearch
  - echo vm.max_map_count=262144 >> /etc/sysctl.conf
  - sysctl -w vm.max_map_count=262144

  # Connect instances to ECS cluster
  - echo ECS_CLUSTER=${aws_ecs_cluster.main.name} >> /etc/ecs/ecs.config

  # Mount EFS volume
  - file_system_id_01=${var.efs_volume}
  - efs_directory=/mnt/efs
  - mkdir -p $${efs_directory}
  - echo "$${file_system_id_01}:/ $${efs_directory} efs tls,_netdev" >> /etc/fstab
  - mount -t efs $${file_system_id_01}:/ $${efs_directory}

  # Create the directories used as container volumes
  - mkdir -p $${efs_directory}/website/oddzy
  - mkdir -p $${efs_directory}/volumes/srv-racing/db-mongo/data
  - mkdir -p $${efs_directory}/volumes/grafana/data
  - chmod -R 777 $${efs_directory}/volumes/grafana/data
  - mkdir -p $${efs_directory}/volumes/elasticsearch/data
  - chmod -R 777 $${efs_directory}/volumes/elasticsearch/data
  - mkdir -p /etc/nginx

  # Read httpasswd file contents from an encrypted KMS secret and save it to the nginx container volume
  - aws ssm get-parameter --name ${var.internal_password_parameter} --with-decryption --output text --region ${var.region} | awk '{print $4}' | sudo tee /etc/nginx/.htpasswd > /dev/null

  # Update the website volume with the latest version from s3
  - rm -rf $${efs_directory}/website/oddzy/*
  - aws s3 cp s3://oddzy/web/dist.zip /tmp/website.zip
  - unzip /tmp/website.zip -d $${efs_directory}/website/oddzy
EOF
}

// Auto-Scaling Group
resource "aws_autoscaling_group" "ecs_autoscaling_group" {
  max_size             = 1
  min_size             = 0
  desired_capacity     = "${var.run_ec2_instance ? 1 : 0}"
  vpc_zone_identifier  = ["${aws_subnet.private.id}"]
  launch_configuration = "${aws_launch_configuration.ecs_launch_configuration.name}"
  target_group_arns    = ["${aws_lb_target_group.public.arn}", "${aws_lb_target_group.private.arn}"]
  health_check_type    = "EC2"
  
  tag {
    key = "name"
    value = "${var.application_name}"
    propagate_at_launch = true
  }
}

// IAM policy / role used by ECS containers

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

// IAM policy / role used by the EC2 instances that run ECS

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

resource "aws_iam_policy" "internal_password" {
  name        = "OddzyInternalPasswordAccess"
  description = "Policy allowing access to the encrypted internal password parameter"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ssm:GetParameter*",
        "kms:Decrypt"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:ssm:${var.region}:${data.aws_caller_identity.current.account_id}:parameter/${var.internal_password_parameter}"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ecs_instance_password_access" {
  role       = "${aws_iam_role.ecs_instance_role.name}"
  policy_arn = "${aws_iam_policy.internal_password.arn}"
}

resource "aws_iam_instance_profile" "ecs_instance_profile" {
  path = "/"
  role = "${aws_iam_role.ecs_instance_role.id}"
}


resource "aws_iam_policy" "website" {
  name        = "OddzyWebsiteAccess"
  description = "Policy allowing access to the zipped website in s3"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::${var.bucket_name}/web/*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ecs_instance_website_access" {
  role       = "${aws_iam_role.ecs_instance_role.name}"
  policy_arn = "${aws_iam_policy.website.arn}"
}
