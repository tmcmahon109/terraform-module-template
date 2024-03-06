resource "aws_security_group" "example" {
  name   = var.sg_name
  tags = {
    Name = var.sg_name
  }
}