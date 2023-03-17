FROM golang:1.19.7-bullseye
SHELL ["/bin/bash", "-c"]
COPY ./authservice /app/authservice
COPY ./common /app/common
WORKDIR /app/authservice
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go install .

ENTRYPOINT exec /go/bin/authservice --rpcPort ${RPC_PORT} --httpPort ${HTTP_PORT} --address ${ADDRESS}
EXPOSE ${RPC_PORT}
EXPOSE ${HTTP_PORT}