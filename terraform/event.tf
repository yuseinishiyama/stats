resource "aws_cloudwatch_event_rule" "this" {
  name                = "stats"
  schedule_expression = "cron(0/5 * * * ? *)"
}

resource "aws_cloudwatch_event_target" "this" {
  rule      = "stats"
  target_id = "stats"
  arn       = aws_ecs_cluster.this.arn
  role_arn  = aws_iam_role.task.arn
  ecs_target {
    launch_type         = "FARGATE"
    platform_version    = "LATEST"
    task_count          = 1
    task_definition_arn = aws_ecs_task_definition.this.arn
    network_configuration {
      subnets = ["subnet-dd3ccb87"]
      security_groups = ["vpc-ed3f4b8b"]
      assign_public_ip = true
    }
  }
}
