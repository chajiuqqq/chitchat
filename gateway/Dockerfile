FROM golang:1.19.7-bullseye
SHELL ["/bin/bash", "-c"]
COPY . /app/gateway
WORKDIR /app/gateway
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go install .
ENV GATEWAY_PORT=7000
ENV CONSUL_PORT=8500
ENV CONSUL_HOST=0.0.0.0
ENTRYPOINT exec /go/bin/gateway --port ${GATEWAY_PORT} --consul-host ${CONSUL_HOST} --consul-port ${CONSUL_PORT}
EXPOSE ${GATEWAY_PORT}