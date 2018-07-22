FROM prom/statsd-exporter:v0.6.0

COPY statsd_mapping.conf /tmp/statsd_mapping.conf

CMD [ "-statsd.mapping-config=/tmp/statsd_mapping.conf" ]