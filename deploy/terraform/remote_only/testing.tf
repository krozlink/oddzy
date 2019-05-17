// EC2 instance deployed to the public subnet - used as a jump box for testing
resource aws_instance testing {
  ami                         = "ami-04c715c762e2de351"
  instance_type               = "t2.medium"
  availability_zone           = "ap-southeast-2a"
  associate_public_ip_address = true
  key_name                    = "${var.ec2_key_pair}"
  subnet_id                   = "${aws_subnet.public-a.id}"
  vpc_security_group_ids      = ["${aws_security_group.testing.id}"]

  count = "${var.run_test_instance ? 1 : 0 }"

  user_data = <<EOF
<powershell>
# Install Chocolatey
iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))

# Globally Auto confirm every action
choco feature enable -n allowGlobalConfirmation

# Install Putty
choco install putty
</powershell>
EOF

  tags {
    Name = "OddzyTesting"
  }
}

// Security group allowing RDP to the jump box
resource aws_security_group testing {
  description = "Used for testing"
  vpc_id      = "${aws_vpc.main.id}"

  ingress {
    from_port   = 3389
    to_port     = 3389
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
resource "aws_ssm_document" "delete_race_data" {
  name          = "${var.application_name}-${var.application_stage}-delete-racing-db"
  document_type = "Command"

  content = <<DOC
  {
    "schemaVersion": "2.2",
    "description": "Delete the mongo DB volume",
    "mainSteps": [
    {
        "action":"aws:runShellScript",
         "name":"deleteDBData",
         "inputs":{
            "runCommand":[
              "rm -rf /mnt/efs/volumes/srv-racing/db-mongo/data/*"
            ]
         }
    }]
  }
DOC
}