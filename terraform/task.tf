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

resource "aws_cloudwatch_event_target" "worker" {
  target_id = "StatsWorker"
  rule      = aws_cloudwatch_event_rule.worker.name
  arn       = aws_ecs_cluster.this.arn
  # use auto-generated role. terraform managed role didn't work
  role_arn = "arn:aws:iam::921647845311:role/ecsEventsRole"
  ecs_target {
    launch_type         = "FARGATE"
    task_count          = 1
    task_definition_arn = aws_ecs_task_definition.worker.arn
    network_configuration {
      subnets          = [data.aws_subnet.container.id]
      security_groups  = [data.aws_security_group.container.id]
      assign_public_ip = false
    }
  }
}

resource "aws_cloudwatch_event_rule" "worker" {
  name                = "stats-worker"
  schedule_expression = "cron(0 0/12 * * ? *)"
}

# bot
resource "aws_ecs_task_definition" "bot" {
  family                   = "bot"
  task_role_arn            = aws_iam_role.task.arn
  execution_role_arn       = aws_iam_role.task_execution.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  container_definitions    = file("containers/bot.json")
}

resource "aws_ecs_service" "bot" {
  name            = "bot"
  cluster         = aws_ecs_cluster.this.id
  task_definition = aws_ecs_task_definition.bot.arn
  launch_type     = "FARGATE"
  desired_count   = 1
  network_configuration {
    subnets          = [data.aws_subnet.container.id]
    security_groups  = [data.aws_security_group.container.id]
    assign_public_ip = true
  }
}
