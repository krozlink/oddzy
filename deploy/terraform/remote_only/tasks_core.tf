// ------- Consul --------------
resource aws_ecs_task_definition consul {
  family                = "consul"
  container_definitions = "${file("task-definitions/core/consul.json")}"
  network_mode          = "host"
}

resource aws_ecs_service consul {
  name            = "consul"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.consul.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Nginx --------------
resource aws_ecs_task_definition nginx {
  family                = "nginx"
  container_definitions = "${file("task-definitions/core/nginx.json")}"
  network_mode          = "host"

  volume {
    name      = "password"
    host_path = "/etc/nginx"
  }

  volume {
    name      = "website"
    host_path = "/mnt/efs/website/oddzy"
  }
}

resource aws_ecs_service nginx {
  name            = "nginx"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.nginx.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Elasticsearch --------------
resource aws_ecs_task_definition elasticsearch {
  family                = "elasticsearch"
  container_definitions = "${file("task-definitions/core/elasticsearch.json")}"
  network_mode          = "host"

  volume {
    name      = "data"
    host_path = "/mnt/efs/volumes/elasticsearch/data"
  }
}

resource aws_ecs_service elasticsearch {
  name            = "elasticsearch"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.elasticsearch.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Kibana --------------
resource aws_ecs_task_definition kibana {
  family                = "kibana"
  container_definitions = "${file("task-definitions/core/kibana.json")}"
  network_mode          = "host"
}

resource aws_ecs_service kibana {
  name            = "kibana"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.kibana.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Logstash --------------
resource aws_ecs_task_definition logstash {
  family                = "logstash"
  container_definitions = "${file("task-definitions/core/logstash.json")}"
  network_mode          = "host"
}

resource aws_ecs_service logstash {
  name            = "logstash"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.logstash.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Micro Web --------------
resource aws_ecs_task_definition micro-web {
  family                = "micro-web"
  container_definitions = "${file("task-definitions/core/micro-web.json")}"
  network_mode          = "host"
}

resource aws_ecs_service micro-web {
  name            = "micro-web"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.micro-web.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Micro API --------------
resource aws_ecs_task_definition micro-api {
  family                = "micro-api"
  container_definitions = "${file("task-definitions/core/micro-api.json")}"
  network_mode          = "host"
}

resource aws_ecs_service micro-api {
  name            = "micro-api"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.micro-api.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Prometheus --------------
resource aws_ecs_task_definition prometheus {
  family                = "prometheus"
  container_definitions = "${file("task-definitions/core/prometheus.json")}"
  network_mode          = "host"
}

resource aws_ecs_service prometheus {
  name            = "prometheus"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.prometheus.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Grafana --------------
resource aws_ecs_task_definition grafana {
  family                = "grafana"
  container_definitions = "${file("task-definitions/core/grafana.json")}"
  network_mode          = "host"

  volume {
    name      = "data"
    host_path = "/mnt/efs/volumes/grafana/data"
  }
}

resource aws_ecs_service grafana {
  name            = "grafana"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.grafana.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Statsd --------------
resource aws_ecs_task_definition statsd {
  family                = "statsd"
  container_definitions = "${file("task-definitions/core/statsd.json")}"
  network_mode          = "host"
}

resource aws_ecs_service statsd {
  name            = "statsd"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.statsd.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- NATS --------------
resource aws_ecs_task_definition nats {
  family                = "nats"
  container_definitions = "${file("task-definitions/core/nats.json")}"
  network_mode          = "host"
}

resource aws_ecs_service nats {
  name            = "nats"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.nats.arn}"
  desired_count   = "${var.run_core_tasks ? 1 : 0}"

  placement_constraints {
    type = "distinctInstance"
  }
}
