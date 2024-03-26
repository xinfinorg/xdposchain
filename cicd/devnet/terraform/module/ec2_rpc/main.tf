variable network {
  type = string
}
variable vpc_id {
  type = string
}
variable aws_subnet_id {
  type = string
}
variable ami_id {
  type = string
}
variable instance_type {
  type = string
}
variable ssh_key_name {
  type = string
}
variable sg_id {
  type = string
}
variable rpc_image {
  type = string
}
variable route53_zone_id {
  type = string
}

resource "aws_instance" "rpc_instance" {
  instance_type           = var.instance_type 
  ami                     = var.ami_id
  tags                    = {
                              Name = var.network
                            }
  key_name                = var.key_name
  vpc_security_group_ids  = [var.sg_id]
  ebs_block_device {
    device_name = "/dev/sda1"
    volume_size = 500
  }


  #below still need to remove git checkout {{branch}} after files merged to master
  user_data = <<-EOF
              #!/bin/bash
              sudo yum update -y
              sudo yum upgrade -y
              sudo yum install git -y
              sudo yum install docker -y
              mkdir -p /root/.docker/cli-plugins
              curl -SL https://github.com/docker/compose/releases/download/v2.25.0/docker-compose-linux-x86_64 -o /root/.docker/cli-plugins/docker-compose
              sudo chmod +x /root/.docker/cli-plugins/docker-compose
              echo checking compose version
              docker compose version
              sudo systemctl enable docker
              sudo systemctl start docker
              mkdir -p /work
              cd /work
              git clone https://github.com/XinFinOrg/XinFin-Node
              cd /work/XinFin-Node/
              git checkout ec2-rpcs
              cd /work/XinFin-Node/${var.network}
              export RPC_IMAGE="${var.rpc_image}"
              echo RPC_IMAGE=$RPC_IMAGE
              ./docker-up.sh
              EOF
}

#TODO: enable after authorized
resource "aws_route53_record" "rpc_route53" {
  zone_id = var.route53_zone_id
  name    = "${var.network}_rpc"
  type    = "A"
  ttl     = "300"
  records = [aws_instance.rpc_instance.public_ip]
}

