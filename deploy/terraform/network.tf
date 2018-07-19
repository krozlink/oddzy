// VPC
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags {
    Name = "${var.application_name}"
  }
}

// Internet gateway
resource "aws_internet_gateway" "gw" {
  vpc_id = "${aws_vpc.main.id}"

  tags {
    Name = "${var.application_name}"
  }
}

resource "aws_nat_gateway" "gw" {
  allocation_id = "${var.nat_eip_allocation_id}"
  subnet_id     = "${aws_subnet.private.id}"

  depends_on = ["aws_internet_gateway.gw"]

  tags {
    Name = "${var.application_name}"
  }
}

// SUBNET - Public
resource "aws_subnet" "public" {
  vpc_id                  = "${aws_vpc.main.id}"
  cidr_block              = "10.0.0.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-southeast-2a"

  tags {
    Name = "${var.application_name}-public-a"
  }
}

// SUBNET - Private
resource "aws_subnet" "private" {
  vpc_id                  = "${aws_vpc.main.id}"
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = false
  availability_zone       = "ap-southeast-2a"

  tags {
    Name = "${var.application_name}-private-a"
  }
}

resource "aws_subnet" "public-b" {
  vpc_id                  = "${aws_vpc.main.id}"
  cidr_block              = "10.0.2.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-southeast-2b"

  tags {
    Name = "${var.application_name}-public-b"
  }
}

// Route table
resource "aws_route_table" "public" {
  vpc_id = "${aws_vpc.main.id}"

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.gw.id}"
  }

  tags {
    Name = "${var.application_name}-public"
  }
}

// Route table
resource "aws_route_table" "private" {
  vpc_id = "${aws_vpc.main.id}"

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = "${aws_nat_gateway.gw.id}"
  }

  tags {
    Name = "${var.application_name}-private"
  }
}

resource "aws_route_table_association" "public" {
  subnet_id      = "${aws_subnet.public.id}"
  route_table_id = "${aws_route_table.public.id}"
}

resource "aws_route_table_association" "private" {
  subnet_id      = "${aws_subnet.private.id}"
  route_table_id = "${aws_route_table.private.id}"
}

resource "aws_route_table_association" "public-b" {
  subnet_id      = "${aws_subnet.public-b.id}"
  route_table_id = "${aws_route_table.public.id}"
}

// Security Groups
resource "aws_security_group" "main_sg" {
  description = "Main VPC security group"
  vpc_id      = "${aws_vpc.main.id}"

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.application_name}"
  }
}
