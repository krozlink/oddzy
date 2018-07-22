resource aws_ecs_task_definition consul {
  family                = "consul"
  container_definitions = "${file("task-definitions/consul.json")}"
  network_mode          = "host"
}

resource aws_ecs_task_definition nginx {
  family                = "nginx"
  container_definitions = "${file("task-definitions/nginx.json")}"
  network_mode          = "host"
}
