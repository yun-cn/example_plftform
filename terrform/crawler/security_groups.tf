resource "aws_security_group" "allow_current_ip" {
  name        = "allow_current_ip"
  description = "Allow inbound traffic from current IP only."

  # vpc_id      = "${aws_vpc.spacemarket_crawler.id}"

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["${data.external.whatsmyip.result["ip"]}/32"]
  }
  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
    prefix_list_ids = []
  }
  tags = {
    Name = "allow_current_ip"
  }
}

data "external" "whatsmyip" {
  program = ["sh", "${path.module}/whatsmyip.sh"]
}

output "current_ip" {
  value = "${data.external.whatsmyip.result["ip"]}"
}
