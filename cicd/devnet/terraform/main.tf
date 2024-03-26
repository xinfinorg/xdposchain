terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.13.1"
    }
  }
}

# Default
provider "aws" {
  region  = "us-east-1"
}

provider "aws" {
  alias = "us-east-2"
  region  = "us-east-2"
}

module "us-east-2" {
  source = "./module/region"
  region = "us-east-2"
  devnetNodeKeys = local.devnetNodeKeys["us-east-2"]
  logLevel = local.logLevel
  devnet_xdc_ecs_tasks_execution_role_arn = aws_iam_role.devnet_xdc_ecs_tasks_execution_role.arn
  docker_tag = var.docker_tag
  providers = {
    aws = aws.us-east-2
  }
}

provider "aws" {
  alias = "eu-west-1"
  region  = "eu-west-1"
}

module "eu-west-1" {
  source = "./module/region"
  region = "eu-west-1"
  devnetNodeKeys = local.devnetNodeKeys["eu-west-1"]
  logLevel = local.logLevel
  devnet_xdc_ecs_tasks_execution_role_arn = aws_iam_role.devnet_xdc_ecs_tasks_execution_role.arn
  docker_tag = var.docker_tag
  providers = {
    aws = aws.eu-west-1
  }
}

provider "aws" {
  alias = "ap-southeast-2"
  region  = "ap-southeast-2"
}

module "ap-southeast-2" {
  source = "./module/region"
  region = "ap-southeast-2"
  devnetNodeKeys = local.devnetNodeKeys["ap-southeast-2"]
  logLevel = local.logLevel
  devnet_xdc_ecs_tasks_execution_role_arn = aws_iam_role.devnet_xdc_ecs_tasks_execution_role.arn
  docker_tag = var.docker_tag
  providers = {
    aws = aws.ap-southeast-2
  }
}

provider "aws" {
  region  = "ap-southeast-1"
}

resource "aws_security_group" "rpc_sg" {
  name_prefix = "rpc-sg"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  
  ingress {
    from_port   = 30303
    to_port     = 30303
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8545
    to_port     = 8545
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8555
    to_port     = 8555
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

module "devnet_rpc" {
  source = "./modules/ec2_rpc"
  network = "devnet"
  vpc_id = local.vpc_id
  aws_subnet_id = local.aws_subnet_id
  ami_id = local.ami_id
  instance_type = "t3.large"
  ssh_key_name = local.ssh_key_name
  sg_id = aws_security_group.rpc_sg.id
  xdpos_image_tag = local.xdpos_image_tag 
  route53_zone_id = local.route53_zone_id
}

module "testnet_rpc" {
  source = "./modules/ec2_rpc"
  network = "testnet"
  vpc_id = local.vpc_id
  aws_subnet_id = local.aws_subnet_id
  ami_id = local.ami_id
  instance_type = "t3.large"
  ssh_key_name = local.ssh_key_name
  sg_id = aws_security_group.rpc_sg.id
  xdpos_image_tag = local.xdpos_image_tag 
  route53_zone_id = local.route53_zone_id
}

module "mainnet_rpc" {
  source = "./modules/ec2_rpc"
  network = "mainnet"
  vpc_id = local.vpc_id
  aws_subnet_id = local.aws_subnet_id
  ami_id = local.ami_id
  instance_type = "t3.large"
  ssh_key_name = local.ssh_key_name
  sg_id = aws_security_group.rpc_sg.id
  xdpos_image_tag = local.xdpos_image_tag 
  route53_zone_id = local.route53_zone_id
}