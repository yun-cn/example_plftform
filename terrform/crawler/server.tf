data "aws_ami" "amazonLinux2AmiTokyo" {
  most_recent = true

  filter {
    name   = "owner-alias"
    values = ["amazon"]
  }

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-2*"]
  }
}

resource "aws_instance" "crawler" {
  ami             = "${data.aws_ami.amazonLinux2AmiTokyo.id}"
  instance_type   = "t2.micro"
  key_name        = "spacemarket_crawler"
  security_groups = ["${aws_security_group.allow_current_ip.id}"]
}

output "crawler_ip" {
  description = "New IP of crawler"
  value       = "${aws_instance.crawler.public_ip}"
}
