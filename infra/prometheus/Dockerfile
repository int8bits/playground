FROM alpine:3

RUN apk update
RUN apk add curl tar gzip

WORKDIR /opt/prom
RUN wget -O prometheus.tar.gz "https://github.com/prometheus/prometheus/releases/download/v2.41.0/prometheus-2.41.0.linux-amd64.tar.gz"
RUN tar xvfz prometheus.tar.gz --strip-components 1
RUN ls -la
RUN more prometheus.yml

EXPOSE 9090

ENTRYPOINT [ "/opt/prom/prometheus", "--config.file=prometheus.yml" ]
