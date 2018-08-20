// ------- srv/racing --------------
resource aws_ecs_task_definition srv-racing {
  family                = "srv-racing"
  container_definitions = "${file("task-definitions/app/srv-racing.json")}"
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
  desired_count   = "${var.run_app_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- srv/race-scraper --------------
resource aws_ecs_task_definition srv-race-scraper {
  family                = "srv-race-scraper"
  container_definitions = "${file("task-definitions/app/srv-race-scraper.json")}"
  network_mode          = "host"
}

resource aws_ecs_service srv-race-scraper {
  name            = "srv-race-scraper"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.srv-race-scraper.arn}"
  desired_count   = "${var.run_app_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- api/racing --------------
resource aws_ecs_task_definition api-racing {
  family                = "api-racing"
  container_definitions = "${file("task-definitions/app/api-racing.json")}"
  network_mode          = "host"
}

resource aws_ecs_service api-racing {
  name            = "api-racing"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.api-racing.arn}"
  desired_count   = "${var.run_app_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}


// ------- web/price-updater --------------
resource aws_ecs_task_definition price-updater {
  family                = "price-updater"
  container_definitions = "${file("task-definitions/app/price-updater.json")}"
  network_mode          = "host"
}

resource aws_ecs_service price-updater {
  name            = "price-updater"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.price-updater.arn}"
  desired_count   = "${var.run_app_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}
