// EC2 instance deployed to the public subnet - used as a jump box for testing
resource aws_instance testing {
  ami                         = "ami-d48623b6"
  instance_type               = "t2.medium"
  availability_zone           = "ap-southeast-2a"
  associate_public_ip_address = true
  key_name                    = "${var.ec2_key_pair}"
  subnet_id                   = "${aws_subnet.public-a.id}"
  vpc_security_group_ids      = ["${aws_security_group.testing.id}"]

  count = "${var.test_instance ? 1 : 0 }"

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

resource aws_security_group testing {
  description = "Used for testing"
  vpc_id      = "${aws_vpc.main.id}"

  # WinRM access from anywhere
  ingress {
    from_port   = 3389
    to_port     = 3389
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # outbound internet access
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
