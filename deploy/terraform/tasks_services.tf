// ------- srv/racing --------------
resource aws_ecs_task_definition srv-racing {
  family                = "srv-racing"
  container_definitions = "${file("task-definitions/services/srv-racing.json")}"
  network_mode          = "host"

  volume {
    name      = "data"
    host_path = "/mnt/efs/volumes/srv-racing/db-mongo/data"
  }
}

resource aws_ecs_service srv-racing {
  name            = "srv-racing"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.srv-racing.arn}"
  desired_count   = 1

  placement_constraints {
    type = "distinctInstance"
  }
}
