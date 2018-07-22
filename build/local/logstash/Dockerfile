FROM docker.elastic.co/logstash/logstash:6.3.0

COPY logstash.conf /etc/logstash/conf.d/logstash.conf

CMD [ "-f","/etc/logstash/conf.d/" ]