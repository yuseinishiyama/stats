# worker
resource "aws_ecs_task_definition" "worker" {
  family                   = "worker"
  task_role_arn            = aws_iam_role.task.arn
  execution_role_arn       = aws_iam_role.task_execution.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  container_definitions    = file("containers/worker.json")
}

# bot
resource "aws_apprunner_service" "bot" {
  service_name = "stats-bot"

  source_configuration {
    authentication_configuration {
      access_role_arn = aws_iam_role.ecr_access.arn
    }
    image_repository {
      image_configuration {
        port          = "80"
        start_command = "bot"
        runtime_environment_variables = {
          FB_VERIFY_TOKEN = var.fb_verify_token
          FB_ACCESS_TOKEN = var.fb_access_token
        }
      }
      image_identifier      = "${aws_ecr_repository.this.repository_url}:amd64"
      image_repository_type = "ECR"
    }
  }
}

variable "fb_verify_token" {
  description = "FB Messenger verification token"
  type        = string
}

variable "fb_access_token" {
  description = "FB Messenger acesss token"
  type        = string
}
