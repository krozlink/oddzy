// ------- Consul --------------
resource aws_ecs_task_definition consul {
  family                = "consul"
  container_definitions = "${file("task-definitions/consul.json")}"
  network_mode          = "host"
}

resource aws_ecs_service consul {
  name            = "consul"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.consul.arn}"
  desired_count   = 1

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Nginx --------------
resource aws_ecs_task_definition nginx {
  family                = "nginx"
  container_definitions = "${file("task-definitions/nginx.json")}"
  network_mode          = "host"
}

resource aws_ecs_service nginx {
  name            = "nginx"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.nginx.arn}"
  desired_count   = 1

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Elasticsearch --------------
resource aws_ecs_task_definition elasticsearch {
  family                = "elasticsearch"
  container_definitions = "${file("task-definitions/elasticsearch.json")}"
  network_mode          = "host"
}

resource aws_ecs_service elasticsearch {
  name            = "elasticsearch"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.elasticsearch.arn}"
  desired_count   = 1

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Kibana --------------
resource aws_ecs_task_definition kibana {
  family                = "kibana"
  container_definitions = "${file("task-definitions/kibana.json")}"
  network_mode          = "host"
}

resource aws_ecs_service kibana {
  name            = "kibana"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.kibana.arn}"
  desired_count   = 1

  placement_constraints {
    type = "distinctInstance"
  }
}

// ------- Logstash --------------
resource aws_ecs_task_definition logstash {
  family                = "logstash"
  container_definitions = "${file("task-definitions/logstash.json")}"
  network_mode          = "host"
}

resource aws_ecs_service logstash {
  name            = "logstash"
  cluster         = "${aws_ecs_cluster.main.name}"
  task_definition = "${aws_ecs_task_definition.logstash.arn}"
  desired_count   = 1

  placement_constraints {
    type = "distinctInstance"
  }
}
