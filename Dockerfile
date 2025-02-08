FROM debian:stable-slim

EXPOSE 9133

WORKDIR /app
COPY ./bin/parse-to-metrics-exporter-amd /app/exporter
RUN chmod a+x /app/exporter
ENTRYPOINT ["/app/exporter"]