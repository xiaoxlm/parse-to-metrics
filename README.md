# 构建二进制
```shell
make build-amd-linux
```

# 构建镜像
```shell
 make build-image
```

# docker 运行
```shell
docker run -d -p 9133:9133 \
 --name  parse-to-metrics-exporter \
 -e AI_METRICS_LABEL=mfu \
 -e NODE_LABEL="your node unique flag" \
 -e LOKI_URL="your loki url" \
parse-to-metrics-exporter:v1.0.0
```
