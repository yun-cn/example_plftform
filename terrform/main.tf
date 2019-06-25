provider "aws" {
  region = "ap-northeast-1"
}

# terraform import aws_key_pair.spacemarket_crawler spacemarket_crawler
resource "aws_key_pair" "spacemarket_crawler" {
  key_name   = "spacemarket_crawler"
  public_key = "${file("${path.module}/spacemarket_crawler.pub")}"
}

module "crawler" {
  source = "./crawler"
}

# terraform output -module=crawler

