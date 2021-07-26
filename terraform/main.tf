terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }

  backend "s3" {
    bucket  = "yuseinishiyama-terraform-remote-state"
    key     = "github.com/yuseinishiyama/stats.tfstate"
    region  = "us-east-1"
    profile = "private"
  }
}

provider "aws" {
  region  = "eu-west-1"
  profile = "private"
}

resource "aws_ecs_cluster" "this" {
  # name "stats" didn't work for unknown reasons
  name = "my-stats"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

resource "aws_ecr_repository" "this" {
  name = "stats"
}
