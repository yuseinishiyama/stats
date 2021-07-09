resource "aws_cloudwatch_event_rule" "this" {
  name                = "stats"
  schedule_expression = "cron(0 0/4 * * ? *)"
}

resource "aws_cloudwatch_event_target" "this" {
  rule      = "stats"
  target_id = "stats"
  arn       = aws_ecs_cluster.this.arn
  # use auto-generated role. terraform managed role didn't work
  role_arn  = "arn:aws:iam::921647845311:role/ecsEventsRole"
  ecs_target {
    launch_type         = "FARGATE"
    platform_version    = "LATEST"
    task_count          = 1
    task_definition_arn = aws_ecs_task_definition.this.arn
    network_configuration {
      subnets = ["subnet-dd3ccb87"]
      security_groups = ["sg-a44cdad9"]
      assign_public_ip = true
    }
  }
}
