// Route53 hosted zone for the domain
data "aws_route53_zone" "main" {
  name         = "${var.domain_name}."
  private_zone = false
}

// Record Set for *.domain.tld
resource aws_route53_record all {
  zone_id = "${data.aws_route53_zone.main.zone_id}"
  name    = "*.${var.domain_name}"
  type    = "A"

  alias {
    name                   = "${aws_lb.main.dns_name}"
    zone_id                = "${aws_lb.main.zone_id}"
    evaluate_target_health = false
  }
}

// Record Set for domain.tld
resource aws_route53_record base {
  zone_id = "${data.aws_route53_zone.main.zone_id}"
  name    = "${var.domain_name}"
  type    = "A"

  alias {
    name                   = "${aws_lb.main.dns_name}"
    zone_id                = "${aws_lb.main.zone_id}"
    evaluate_target_health = false
  }
}
